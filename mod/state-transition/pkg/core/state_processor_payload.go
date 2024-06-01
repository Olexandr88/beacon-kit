// SPDX-License-Identifier: MIT
//
// Copyright (c) 2024 Berachain Foundation
//
// Permission is hereby granted, free of charge, to any person
// obtaining a copy of this software and associated documentation
// files (the "Software"), to deal in the Software without
// restriction, including without limitation the rights to use,
// copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following
// conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// OTHER DEALINGS IN THE SOFTWARE.

package core

import (
	"context"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/errors"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/ssz"
	"golang.org/x/sync/errgroup"
)

// processExecutionPayload processes the execution payload and ensures it
// matches the local state.
func (sp *StateProcessor[
	BeaconBlockT, BeaconBlockBodyT, BeaconBlockHeaderT,
	BeaconStateT, BlobSidecarsT, ContextT,
	DepositT, ExecutionPayloadT, ExecutionPayloadHeaderT,
	ForkDataT, ValidatorT, WithdrawalT, WithdrawalCredentialsT,
]) processExecutionPayload(
	ctx ContextT,
	st BeaconStateT,
	blk BeaconBlockT,
) error {
	var (
		body            = blk.GetBody()
		payload         = body.GetExecutionPayload()
		txsRoot         primitives.Root
		withdrawalsRoot primitives.Root
		g, gCtx         = errgroup.WithContext(context.Background())
	)

	// Skip payload verification if the context is configured as such.
	if !ctx.GetSkipPayloadVerification() {
		g.Go(func() error {
			return sp.validateExecutionPayload(
				gCtx, st, blk, ctx.GetOptimisticEngine(),
			)
		})
	}

	g.Go(func() error {
		var txsRootErr error
		txsRoot, txsRootErr = engineprimitives.Transactions(
			payload.GetTransactions(),
		).HashTreeRoot()
		return txsRootErr
	})

	g.Go(func() error {
		var withdrawalsRootErr error
		withdrawalsRoot, withdrawalsRootErr =
			ssz.MerkleizeListComposite[any, math.U64](
				payload.GetWithdrawals(),
				sp.cs.MaxWithdrawalsPerPayload(),
			)
		return withdrawalsRootErr
	})

	// If deriving either of the roots or verifying the payload fails, return
	// the error.
	if err := g.Wait(); err != nil {
		return err
	}

	// Set the latest execution payload header.
	return st.SetLatestExecutionPayloadHeader(
		&types.ExecutionPayloadHeader{
			ExecutionPayloadHeader: &types.ExecutionPayloadHeaderDeneb{
				ParentHash:       payload.GetParentHash(),
				FeeRecipient:     payload.GetFeeRecipient(),
				StateRoot:        payload.GetStateRoot(),
				ReceiptsRoot:     payload.GetReceiptsRoot(),
				LogsBloom:        payload.GetLogsBloom(),
				Random:           payload.GetPrevRandao(),
				Number:           payload.GetNumber(),
				GasLimit:         payload.GetGasLimit(),
				GasUsed:          payload.GetGasUsed(),
				Timestamp:        payload.GetTimestamp(),
				ExtraData:        payload.GetExtraData(),
				BaseFeePerGas:    payload.GetBaseFeePerGas(),
				BlockHash:        payload.GetBlockHash(),
				TransactionsRoot: txsRoot,
				WithdrawalsRoot:  withdrawalsRoot,
				BlobGasUsed:      payload.GetBlobGasUsed(),
				ExcessBlobGas:    payload.GetExcessBlobGas(),
			},
		},
	)
}

// validateExecutionPayload validates the execution payload against both local
// state
// and the execution engine.
func (sp *StateProcessor[
	BeaconBlockT, BeaconBlockBodyT, BeaconBlockHeaderT,
	BeaconStateT, BlobSidecarsT, ContextT,
	DepositT, ExecutionPayloadT, ExecutionPayloadHeaderT,
	ForkDataT, ValidatorT, WithdrawalT, WithdrawalCredentialsT,
]) validateExecutionPayload(
	ctx context.Context,
	st BeaconStateT,
	blk BeaconBlockT,
	optimisticEngine bool,
) error {
	body := blk.GetBody()
	payload := body.GetExecutionPayload()

	lph, err := st.GetLatestExecutionPayloadHeader()
	if err != nil {
		return err
	}

	// We want to check to ensure the chain is canonical with respect to the
	// parent hash before we let the execution client know about the
	// payload,
	// this is to prevent Polygon style re-orgs from being triggered by a
	// malicious actor who tries to force clients to accept a non-canonical
	// block that passes block validity checks.
	if safeHash := lph.GetBlockHash(); safeHash != payload.GetParentHash() {
		return errors.Wrapf(
			ErrParentRootMismatch,
			"parent block with hash %x is not finalized, expected finalized hash %x",
			payload.GetParentHash(),
			safeHash,
		)
	}

	parentBeaconBlockRoot := blk.GetParentBlockRoot()
	if err = sp.executionEngine.VerifyAndNotifyNewPayload(
		ctx, engineprimitives.BuildNewPayloadRequest(
			payload,
			body.GetBlobKzgCommitments().ToVersionedHashes(),
			&parentBeaconBlockRoot,
			optimisticEngine,
		),
	); err != nil {
		return err
	}

	// Get the current epoch.
	slot, err := st.GetSlot()
	if err != nil {
		return err
	}

	// When we are verifying a payload we expect that it was produced by
	// the proposer for the slot that it is for.
	expectedMix, err := st.GetRandaoMixAtIndex(
		uint64(sp.cs.SlotToEpoch(slot)) % sp.cs.EpochsPerHistoricalVector())
	if err != nil {
		return err
	}

	// Ensure the prev randao matches the local state.
	if payload.GetPrevRandao() != expectedMix {
		return errors.Wrapf(
			ErrRandaoMixMismatch,
			"prev randao does not match, expected: %x, got: %x",
			expectedMix, payload.GetPrevRandao(),
		)
	}

	// TODO: Verify timestamp data once Clock is done.
	// if expectedTime, err := spec.TimeAtSlot(slot, genesisTime); err !=
	// nil { 	return errors.Newf("slot or genesis time in state is corrupt,
	// cannot
	// compute time: %v", err)
	// } else if payload.Timestamp != expectedTime {
	// 	return errors.Newf("state at slot %d, genesis time %d, expected
	// execution
	// payload time %d, but got %d",
	// 		slot, genesisTime, expectedTime, payload.Timestamp)
	// }

	// Verify the number of blobs.
	blobKzgCommitments := body.GetBlobKzgCommitments()
	if uint64(len(blobKzgCommitments)) > sp.cs.MaxBlobsPerBlock() {
		return errors.Wrapf(
			ErrExceedsBlockBlobLimit,
			"expected: %d, got: %d",
			sp.cs.MaxBlobsPerBlock(), len(blobKzgCommitments),
		)
	}

	// Verify the number of withdrawals.
	// TODO: This is in the wrong spot I think.
	if withdrawals := payload.GetWithdrawals(); uint64(
		len(payload.GetWithdrawals()),
	) > sp.cs.MaxWithdrawalsPerPayload() {
		return errors.Newf(
			"too many withdrawals, expected: %d, got: %d",
			sp.cs.MaxWithdrawalsPerPayload(), len(withdrawals),
		)
	}
	return nil
}
