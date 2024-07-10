// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2024, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package store

import (
	"context"

	// "github.com/berachain/beacon-kit/mod/da/pkg/types"
	"github.com/berachain/beacon-kit/mod/errors"
	types "github.com/berachain/beacon-kit/mod/interfaces/pkg/consensus-types"
	datypes "github.com/berachain/beacon-kit/mod/interfaces/pkg/da/types"
	db "github.com/berachain/beacon-kit/mod/interfaces/pkg/storage"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/sourcegraph/conc/iter"
)

// Store is the default implementation of the AvailabilityStore.
type Store[
	BeaconBlockBodyT types.BeaconBlockBody[
		BeaconBlockBodyT, DepositT, Eth1DataT, ExecutionPayloadT,
	],
	BeaconBlockHeaderT any,
	BlobSidecarT datypes.BlobSidecar[
		BlobSidecarT, BeaconBlockHeaderT,
	],
	BlobSidecarsT datypes.BlobSidecars[
		BlobSidecarsT, BlobSidecarT,
	],
	DepositT any,
	Eth1DataT any,
	ExecutionPayloadT any,
] struct {
	// IndexDB is a basic database interface.
	db.IndexDB
	// logger is used for logging.
	logger log.Logger[any]
	// chainSpec contains the chain specification.
	chainSpec common.ChainSpec
}

// New creates a new instance of the AvailabilityStore.
func New[
	BeaconBlockBodyT types.BeaconBlockBody[
		BeaconBlockBodyT, DepositT, Eth1DataT, ExecutionPayloadT,
	],
	BeaconBlockHeaderT any,
	BlobSidecarT datypes.BlobSidecar[
		BlobSidecarT, BeaconBlockHeaderT,
	],
	BlobSidecarsT datypes.BlobSidecars[
		BlobSidecarsT, BlobSidecarT,
	],
	DepositT any,
	Eth1DataT any,
	ExecutionPayloadT any,
](
	db db.IndexDB,
	logger log.Logger[any],
	chainSpec common.ChainSpec,
) *Store[
	BeaconBlockBodyT, BeaconBlockHeaderT, BlobSidecarT,
	BlobSidecarsT, DepositT, Eth1DataT, ExecutionPayloadT,
] {
	return &Store[
		BeaconBlockBodyT, BeaconBlockHeaderT, BlobSidecarT,
		BlobSidecarsT, DepositT, Eth1DataT, ExecutionPayloadT,
	]{
		IndexDB:   db,
		chainSpec: chainSpec,
		logger:    logger,
	}
}

// IsDataAvailable ensures that all blobs referenced in the block are
// stored before it returns without an error.
func (s *Store[BeaconBlockBodyT, _, _, _, _, _, _]) IsDataAvailable(
	_ context.Context,
	slot math.Slot,
	body BeaconBlockBodyT,
) bool {
	for _, commitment := range body.GetBlobKzgCommitments() {
		// Check if the block data is available in the IndexDB
		blockData, err := s.IndexDB.Has(uint64(slot), commitment[:])
		if err != nil || !blockData {
			return false
		}
	}
	return true
}

// Persist ensures the sidecar data remains accessible, utilizing parallel
// processing for efficiency.
func (s *Store[_, _, BlobSidecarT, BlobSidecarsT, _, _, _]) Persist(
	slot math.Slot,
	sidecars BlobSidecarsT,
) error {
	// Exit early if there are no sidecars to store.
	if sidecars.IsNil() || sidecars.Len() == 0 {
		return nil
	}
	sidecar, _ := sidecars.Get(0)

	// Check to see if we are required to store the sidecar anymore, if
	// this sidecar is from outside the required DA period, we can skip it.
	if !s.chainSpec.WithinDAPeriod(
		// slot in which the sidecar was included.
		// (Safe to assume all sidecars are in same slot at this point).
		sidecar.GetSlot(),
		// current slot
		slot,
	) {
		return nil
	}

	// Store each sidecar in parallel.
	if err := errors.Join(iter.Map(
		sidecars.GetSidecars(),
		func(sidecar *BlobSidecarT) error {
			sc := *sidecar
			if (sc).IsNil() {
				return ErrAttemptedToStoreNilSidecar
			}
			bz, err := sc.MarshalSSZ()
			if err != nil {
				return err
			}
			commitment := sc.GetCommitment()
			return s.Set(uint64(slot), commitment[:], bz)
		},
	)...); err != nil {
		return err
	}

	s.logger.Info("Successfully stored all blob sidecars 🚗",
		"slot", slot.Base10(), "num_sidecars", sidecars.Len(),
	)
	return nil
}
