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

package runtime

import (
	"context"

	"github.com/berachain/beacon-kit/mod/beacon/blockchain"
	"github.com/berachain/beacon-kit/mod/consensus-types/pkg/types"
	engineprimitives "github.com/berachain/beacon-kit/mod/engine-primitives/pkg/engine-primitives"
	"github.com/berachain/beacon-kit/mod/log"
	"github.com/berachain/beacon-kit/mod/primitives"
	"github.com/berachain/beacon-kit/mod/primitives/pkg/math"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/runtime/middleware"
	"github.com/berachain/beacon-kit/mod/runtime/pkg/service"
	"github.com/berachain/beacon-kit/mod/state-transition/pkg/core"
)

type BeaconState = core.BeaconState[
	*types.BeaconBlockHeader,
	*types.Eth1Data,
	*types.ExecutionPayloadHeader,
	*types.Fork,
	*types.Validator,
	*engineprimitives.Withdrawal,
]

// BeaconKitRuntime is a struct that holds the
// service registry.
type BeaconKitRuntime[
	AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT],
	BeaconBlockT interface {
		types.RawBeaconBlock[BeaconBlockBodyT, *types.ExecutionPayload]
		NewFromSSZ([]byte, uint32) (BeaconBlockT, error)
		NewWithVersion(
			math.Slot,
			math.ValidatorIndex,
			primitives.Root,
			uint32,
		) (BeaconBlockT, error)
		Empty(uint32) BeaconBlockT
	},
	BeaconBlockBodyT types.RawBeaconBlockBody[*types.ExecutionPayload],
	BeaconStateT core.BeaconState[
		*types.BeaconBlockHeader,
		*types.Eth1Data,
		*types.ExecutionPayloadHeader,
		*types.Fork,
		*types.Validator,
		*engineprimitives.Withdrawal,
	],
	BlobSidecarsT BlobSidecars,
	DepositStoreT DepositStore,
	ExecutionPayloadHeaderT *types.ExecutionPayloadHeader,
	StorageBackendT StorageBackend[
		AvailabilityStoreT, BeaconBlockBodyT,
		BeaconStateT, BlobSidecarsT, DepositStoreT,
	],
] struct {
	// logger is used for logging within the BeaconKitRuntime.
	logger log.Logger[any]
	// services is a registry of services used by the BeaconKitRuntime.
	services *service.Registry
	// storageBackend is the backend storage interface used by the
	// BeaconKitRuntime.
	storageBackend StorageBackendT
	// chainSpec defines the chain specifications for the BeaconKitRuntime.
	chainSpec primitives.ChainSpec
	// abciFinalizeBlockMiddleware handles ABCI interactions for the
	// BeaconKitRuntime.
	abciFinalizeBlockMiddleware *middleware.FinalizeBlockMiddleware[
		BeaconBlockT, BeaconStateT, BlobSidecarsT, ExecutionPayloadHeaderT,
	]
	// abciValidatorMiddleware is responsible for forward ABCI requests to the
	// validator service.
	abciValidatorMiddleware *middleware.ValidatorMiddleware[
		AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT,
		BeaconStateT, BlobSidecarsT, ExecutionPayloadHeaderT, StorageBackendT,
	]
}

// NewBeaconKitRuntime creates a new BeaconKitRuntime
// and applies the provided options.
func NewBeaconKitRuntime[
	AvailabilityStoreT AvailabilityStore[BeaconBlockBodyT, BlobSidecarsT],
	BeaconBlockT interface {
		types.RawBeaconBlock[BeaconBlockBodyT, *types.ExecutionPayload]
		NewFromSSZ([]byte, uint32) (BeaconBlockT, error)
		NewWithVersion(
			math.Slot,
			math.ValidatorIndex,
			primitives.Root,
			uint32,
		) (BeaconBlockT, error)
		Empty(uint32) BeaconBlockT
	},
	BeaconBlockBodyT types.RawBeaconBlockBody[*types.ExecutionPayload],
	BeaconStateT core.BeaconState[
		*types.BeaconBlockHeader, *types.Eth1Data,
		*types.ExecutionPayloadHeader, *types.Fork,
		*types.Validator, *engineprimitives.Withdrawal,
	],
	BlobSidecarsT BlobSidecars,
	DepositStoreT DepositStore,
	ExecutionPayloadHeaderT *types.ExecutionPayloadHeader,
	StorageBackendT blockchain.StorageBackend[
		AvailabilityStoreT,
		BeaconBlockBodyT,
		BeaconStateT,
		BlobSidecarsT,
		*types.Deposit,
		DepositStoreT,
	],
](
	chainSpec primitives.ChainSpec,
	finalizeBlockMiddleware *middleware.FinalizeBlockMiddleware[
		BeaconBlockT, BeaconStateT, BlobSidecarsT, ExecutionPayloadHeaderT,
	],
	logger log.Logger[any],
	services *service.Registry,
	storageBackend StorageBackendT,
	validatorMiddleware *middleware.ValidatorMiddleware[
		AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT,
		BeaconStateT, BlobSidecarsT, ExecutionPayloadHeaderT, StorageBackendT,
	],
) (*BeaconKitRuntime[
	AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT, BeaconStateT,
	BlobSidecarsT, DepositStoreT, ExecutionPayloadHeaderT, StorageBackendT,
], error) {
	return &BeaconKitRuntime[
		AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT, BeaconStateT,
		BlobSidecarsT, DepositStoreT, ExecutionPayloadHeaderT, StorageBackendT,
	]{
		abciFinalizeBlockMiddleware: finalizeBlockMiddleware,
		abciValidatorMiddleware:     validatorMiddleware,
		chainSpec:                   chainSpec,
		logger:                      logger,
		services:                    services,
		storageBackend:              storageBackend,
	}, nil
}

// StartServices starts the services.
func (r *BeaconKitRuntime[
	AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT, BeaconStateT,
	BlobSidecarsT, DepositStoreT, ExecutionPayloadHeaderT, StorageBackendT,
]) StartServices(
	ctx context.Context,
) error {
	return r.services.StartAll(ctx)
}

// ABCIFinalizeBlockMiddleware returns the ABCI handler.
func (r *BeaconKitRuntime[
	AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT, BeaconStateT,
	BlobSidecarsT, DepositStoreT, ExecutionPayloadHeaderT, StorageBackendT,
]) ABCIFinalizeBlockMiddleware() *middleware.FinalizeBlockMiddleware[
	BeaconBlockT, BeaconStateT, BlobSidecarsT, ExecutionPayloadHeaderT,
] {
	return r.abciFinalizeBlockMiddleware
}

// ABCIValidatorMiddleware returns the ABCI validator middleware.
func (r *BeaconKitRuntime[
	AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT, BeaconStateT,
	BlobSidecarsT, DepositStoreT, ExecutionPayloadHeaderT, StorageBackendT,
]) ABCIValidatorMiddleware() *middleware.ValidatorMiddleware[
	AvailabilityStoreT, BeaconBlockT, BeaconBlockBodyT,
	BeaconStateT, BlobSidecarsT, ExecutionPayloadHeaderT, StorageBackendT,
] {
	return r.abciValidatorMiddleware
}
