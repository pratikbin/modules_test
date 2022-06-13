// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package revoke

import (
	baseHelpers "github.com/AssetMantle/modules/schema/helpers/base"
	"github.com/AssetMantle/modules/schema/helpers/constants"
)

var Transaction = baseHelpers.NewTransaction(
	"revoke",
	"",
	"",

	requestPrototype,
	messagePrototype,
	keeperPrototype,

	constants.FromID,
	constants.ToID,
	constants.ClassificationID,
)
