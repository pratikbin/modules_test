// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package properties

import (
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

// TODO add update method
type Property interface {
	GetID() ids.PropertyID
	GetDataID() ids.DataID
	GetKey() ids.ID
	GetType() ids.ID
	GetHash() ids.ID

	traits.Listable
}
