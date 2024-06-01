package backend

import (
	"context"

	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	"github.com/berachain/beacon-kit/mod/node-api/backend/mocks"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/crypto"
	"github.com/stretchr/testify/mock"
)

func NewMockBackend() *Backend {
	sdb := &mocks.StateDB{}
	b := New(func(context.Context, string) StateDB {
		return sdb
	})
	setReturnValues(sdb)
	return b
}

func setReturnValues(sdb *mocks.StateDB) {
	sdb.EXPECT().GetGenesisValidatorsRoot().Return(primitives.Root{0x01}, nil)
	sdb.EXPECT().GetSlot().Return(1, nil)
	sdb.EXPECT().GetLatestExecutionPayloadHeader().Return(nil, nil)
	sdb.EXPECT().SetLatestExecutionPayloadHeader(mock.Anything).Return(nil)
	sdb.EXPECT().GetEth1DepositIndex().Return(0, nil)
	sdb.EXPECT().SetEth1DepositIndex(mock.Anything).Return(nil)
	sdb.EXPECT().GetBalance(mock.Anything).Return(1, nil)
	sdb.EXPECT().SetBalance(mock.Anything, mock.Anything).Return(nil)
	sdb.EXPECT().SetSlot(mock.Anything).Return(nil)
	sdb.EXPECT().GetFork().Return(nil, nil)
	sdb.EXPECT().SetFork(mock.Anything).Return(nil)
	sdb.EXPECT().GetLatestBlockHeader().Return(nil, nil)
	sdb.EXPECT().SetLatestBlockHeader(mock.Anything).Return(nil)
	sdb.EXPECT().
		GetBlockRootAtIndex(mock.Anything).
		Return(primitives.Root{0x01}, nil)
	sdb.EXPECT().
		StateRootAtIndex(mock.Anything).
		Return(primitives.Root{0x01}, nil)
	sdb.EXPECT().GetEth1Data().Return(nil, nil)
	sdb.EXPECT().SetEth1Data(mock.Anything).Return(nil)
	sdb.EXPECT().GetValidators().Return(nil, nil)
	sdb.EXPECT().GetBalances().Return(nil, nil)
	sdb.EXPECT().GetNextWithdrawalIndex().Return(0, nil)
	sdb.EXPECT().SetNextWithdrawalIndex(mock.Anything).Return(nil)
	sdb.EXPECT().GetNextWithdrawalValidatorIndex().Return(0, nil)
	sdb.EXPECT().SetNextWithdrawalValidatorIndex(mock.Anything).Return(nil)
	sdb.EXPECT().GetTotalSlashing().Return(0, nil)
	sdb.EXPECT().SetTotalSlashing(mock.Anything).Return(nil)
	sdb.EXPECT().
		GetRandaoMixAtIndex(mock.Anything).
		Return(primitives.Bytes32{0x01}, nil)
	sdb.EXPECT().GetSlashings().Return(nil, nil)
	sdb.EXPECT().SetSlashingAtIndex(mock.Anything, mock.Anything).Return(nil)
	sdb.EXPECT().GetSlashingAtIndex(mock.Anything).Return(0, nil)
	sdb.EXPECT().GetTotalValidators().Return(0, nil)
	sdb.EXPECT().GetTotalActiveBalances(mock.Anything).Return(0, nil)
	sdb.EXPECT().ValidatorByIndex(mock.Anything).Return(&types.Validator{
		Pubkey:                     crypto.BLSPubkey{0x01},
		WithdrawalCredentials:      types.WithdrawalCredentials{0x01},
		EffectiveBalance:           0,
		Slashed:                    false,
		ActivationEligibilityEpoch: 0,
		ActivationEpoch:            0,
		ExitEpoch:                  0,
		WithdrawableEpoch:          0,
	}, nil)
	sdb.EXPECT().
		UpdateBlockRootAtIndex(mock.Anything, mock.Anything).
		Return(nil)
	sdb.EXPECT().
		UpdateStateRootAtIndex(mock.Anything, mock.Anything).
		Return(nil)
	sdb.EXPECT().
		UpdateRandaoMixAtIndex(mock.Anything, mock.Anything).
		Return(nil)
	sdb.EXPECT().
		UpdateValidatorAtIndex(mock.Anything, mock.Anything).
		Return(nil)
	sdb.EXPECT().ValidatorIndexByPubkey(mock.Anything).Return(0, nil)
	sdb.EXPECT().AddValidator(mock.Anything).Return(nil)
	sdb.EXPECT().GetValidatorsByEffectiveBalance().Return(nil, nil)
}
