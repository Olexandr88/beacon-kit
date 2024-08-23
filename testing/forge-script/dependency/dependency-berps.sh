#!/bin/sh
# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2024, Berachain Foundation. All rights reserved.
# Use of this software is governed by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.
apk update && apk add --no-cache nodejs npm

npm --version
npm install -g bun

bun --version

cd /app/contracts && bun install

echo "Bun installation complete!"

echo "Installing dependencies for Berps"

cd /app/dependency && sh populate-envrc.sh

cp .envrc /app/contracts/script/berps/.envrc

cd /app/contracts/script/berps

forge build

sh deploy-berps-deployer.sh >> output.json
sh deploy-contracts.sh >> output.json

cp output.json /app/contracts/output.json