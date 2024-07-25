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

package beacon

import (
	beacontypes "github.com/berachain/beacon-kit/mod/node-api/handlers/beacon/types"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/common"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
)

// Backend is the interface for backend of the beacon API.
type Backend[
	BlockHeaderT BlockHeader,
	ForkT any,
	ValidatorT any,
] interface {
	GenesisBackend
	BlockBackend[BlockHeaderT]
	RandaoBackend
	ValidatorBackend[ValidatorT]
	HistoricalBackend[ForkT]
}

type BlockHeader interface {
	HashTreeRoot() ([32]byte, error)
	GetSlot() math.Slot
	GetProposerIndex() math.ValidatorIndex
	GetParentBlockRoot() common.Root
	GetStateRoot() common.Root
	GetBodyRoot() common.Root
}

type GenesisBackend interface {
	GenesisValidatorsRoot(
		slot uint64,
	) (common.Root, error)
}

type HistoricalBackend[ForkT any] interface {
	StateRootAtSlot(
		slot uint64,
	) (common.Root, error)
	StateForkAtSlot(
		slot uint64,
	) (ForkT, error)
}

type RandaoBackend interface {
	RandaoAtEpoch(
		slot, epoch uint64,
	) (common.Bytes32, error)
}

type BlockBackend[BeaconBlockHeaderT any] interface {
	BlockRootAtSlot(
		slot uint64,
	) (common.Root, error)
	BlockRewardsAtSlot(
		slot uint64,
	) (*beacontypes.BlockRewardsData, error)
	BlockHeaderAtSlot(
		slot uint64,
	) (BeaconBlockHeaderT, error)
}

type ValidatorBackend[ValidatorT any] interface {
	ValidatorByID(
		slot uint64,
		id string,
	) (*beacontypes.ValidatorData[ValidatorT], error)
	ValidatorsByIDs(
		slot uint64,
		ids []string,
		statuses []string,
	) ([]*beacontypes.ValidatorData[ValidatorT], error)
	ValidatorBalancesByIDs(
		slot uint64,
		ids []string,
	) ([]*beacontypes.ValidatorBalanceData, error)
}