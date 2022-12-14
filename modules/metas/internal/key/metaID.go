// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/metas/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type metaID struct {
	TypeID ids.ID `json:"typeID" valid:"required~required field typeID missing"`
	HashID ids.ID `json:"hashID" valid:"required~required field hashID missing"`
}

var _ ids.ID = (*metaID)(nil)
var _ helpers.Key = (*metaID)(nil)

func (metaID metaID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, metaID.TypeID.Bytes()...)
	Bytes = append(Bytes, metaID.HashID.Bytes()...)

	return Bytes
}
func (metaID metaID) String() string {
	var values []string
	values = append(values, metaID.TypeID.String())
	values = append(values, metaID.HashID.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (metaID metaID) Compare(listable traits.Listable) int {
	return bytes.Compare(metaID.Bytes(), metaIDFromInterface(listable).Bytes())
}
func (metaID metaID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(metaID.Bytes())
}
func (metaID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, metaID{})
}
func (metaID metaID) IsPartial() bool {
	return len(metaID.HashID.Bytes()) == 0
}
func (metaID metaID) Equals(key helpers.Key) bool {
	return metaID.Compare(metaIDFromInterface(key)) == 0
}

func NewMetaID(typeID ids.ID, hashID ids.ID) ids.ID {
	return metaID{
		TypeID: typeID,
		HashID: hashID,
	}
}
