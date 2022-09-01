// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"bytes"

	errorConstants "github.com/AssetMantle/modules/schema/errors/constants"
	"github.com/AssetMantle/modules/schema/ids"
	stringUtilities "github.com/AssetMantle/modules/schema/ids/utilities"
	"github.com/AssetMantle/modules/schema/qualified"
	"github.com/AssetMantle/modules/schema/traits"
)

type assetID struct {
	ids.ClassificationID
	ids.HashID
}

func (assetID assetID) IsOwnableID() {
	// TODO implement me
	panic("implement me")
}

func (assetID assetID) IsAssetID() {
	// TODO implement me
	panic("implement me")
}

var _ ids.AssetID = (*assetID)(nil)

func (assetID assetID) String() string {
	return stringUtilities.JoinIDStrings(assetID.ClassificationID.String(), assetID.HashID.String())
}
func (assetID assetID) Bytes() []byte {
	var Bytes []byte
	Bytes = append(Bytes, assetID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, assetID.HashID.Bytes()...)

	return Bytes
}
func (assetID assetID) Compare(listable traits.Listable) int {
	return bytes.Compare(assetID.Bytes(), assetIDFromInterface(listable).Bytes())
}
func assetIDFromInterface(i interface{}) assetID {
	switch value := i.(type) {
	case assetID:
		return value
	default:
		panic(errorConstants.MetaDataError)
	}
}
func NewAssetID(classificationID ids.ClassificationID, immutables qualified.Immutables) ids.AssetID {
	return assetID{
		ClassificationID: classificationID,
		HashID:           immutables.GenerateHashID(),
	}
}

func PrototypeAssetID() ids.AssetID {
	return assetID{
		ClassificationID: PrototypeClassificationID(),
		HashID:           PrototypeHashID(),
	}
}

func ReadAssetID(assetIDString string) (ids.AssetID, error) {
	if splitAssetIDString := stringUtilities.SplitCompositeIDString(assetIDString); len(splitAssetIDString) == 2 {
		if classificationID, err := ReadClassificationID(splitAssetIDString[0]); err == nil {
			if hashID, err := ReadHashID(splitAssetIDString[1]); err == nil {
				return assetID{
					ClassificationID: classificationID,
					HashID:           hashID,
				}, nil
			}
		}
	}

	if assetIDString == "" {
		return PrototypeAssetID(), nil
	}

	return assetID{}, errorConstants.MetaDataError
}
