// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"
	"strings"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/schema/data"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/traits"
)

type dataID struct {
	Type ids.ID
	Hash ids.ID
}

var _ ids.DataID = (*dataID)(nil)

func (dataID dataID) String() string {
	var values []string
	values = append(values, dataID.Type.String())
	values = append(values, dataID.Hash.String())

	return strings.Join(values, constants.FirstOrderCompositeIDSeparator)
}
func (dataID dataID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, dataID.Type.Bytes()...)
	Bytes = append(Bytes, dataID.Hash.Bytes()...)

	return Bytes
}
func (dataID dataID) Compare(listable traits.Listable) int {
	if compareDataID, err := dataIDFromInterface(listable); err != nil {
		panic(errors.MetaDataError)
	} else {
		return bytes.Compare(dataID.Bytes(), compareDataID.Bytes())
	}
}
func (dataID dataID) GetHash() ids.ID {
	return dataID.Hash
}
func dataIDFromInterface(i interface{}) (dataID, error) {
	switch value := i.(type) {
	case dataID:
		return value, nil
	default:
		return dataID{}, errors.MetaDataError
	}
}
func NewDataID(data data.Data) ids.DataID {
	if data == nil {
		panic(errors.MetaDataError)
	}

	return dataID{
		Type: data.GetType(),
		Hash: data.GenerateHash(),
	}
}
