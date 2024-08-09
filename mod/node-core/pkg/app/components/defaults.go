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

package components

func DefaultComponentsWithStandardTypes[
	AT AI,
]() []any {
	components := []any{
		ProvideAttributesFactory,
		ProvideAvailabilityPruner,
		ProvideAvailibilityStore,
		ProvideBeaconDepositContract,
		ProvideBeaconStore,
		ProvideBlockPruner,
		ProvideBlockStore,
		ProvideBlockStoreService,
		ProvideBlsSigner,
		ProvideBlobProcessor,
		ProvideBlobProofVerifier,
		ProvideBlobVerifier,
		ProvideChainService,
		ProvideChainSpec,
		ProvideDAService,
		ProvideDBManager,
		ProvideDepositPruner,
		ProvideDepositService,
		ProvideDepositStore,
		ProvideEngineClient,
		ProvideExecutionEngine,
		ProvideJWTSecret,
		ProvideLocalBuilder,
		ProvideReportingService,
		ProvideRootStore,
		ProvideRuntimeApp,
		ProvideSDKLogger,
		ProvideServiceRegistry,
		ProvideSidecarFactory,
		ProvideStateProcessor,
		ProvideStorageBackend,
		ProvideStoreOptions,
		ProvideTelemetrySink,
		ProvideTrustedSetup,
		ProvideValidatorService,
		// POC
		ProvideA,
		ProvideB[AT],
	}
	// components = append(components, DefaultNodeAPIComponents()...)
	// components = append(components, DefaultNodeAPIHandlers()...)
	components = append(components, DefaultBrokerProviders()...)
	return components
}

// POC with generics

type A struct{}

func (a *A) Hello() {}

type AI interface{ Hello() }

type B struct{ A AI }

func ProvideA() *A { return &A{} }

func ProvideB[AT AI](a AT) *B { return &B{A: a} }