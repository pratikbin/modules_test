// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants"
	"github.com/AssetMantle/modules/modules/orders/internal/module"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/ids"
	"github.com/AssetMantle/modules/schema/lists"
	baseQualified "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/traits"
	codecUtilities "github.com/AssetMantle/modules/utilities/codec"
)

type orderID struct {
	ClassificationID ids.ID `json:"classificationID"`
	MakerOwnableID   ids.ID `json:"makerOwnableID"`
	TakerOwnableID   ids.ID `json:"takerOwnableID"`
	RateID           ids.ID `json:"rateID"`
	CreationID       ids.ID `json:"creationID"`
	MakerID          ids.ID `json:"makerID"`
	HashID           ids.ID `json:"hashID"`
}

var _ ids.ID = (*orderID)(nil)
var _ helpers.Key = (*orderID)(nil)

func (orderID orderID) Bytes() []byte {
	var Bytes []byte

	rateIDBytes, err := orderID.getRateIDBytes()
	if err != nil {
		return Bytes
	}

	creationIDBytes, err := orderID.getCreationHeightBytes()
	if err != nil {
		return Bytes
	}

	Bytes = append(Bytes, orderID.ClassificationID.Bytes()...)
	Bytes = append(Bytes, orderID.MakerOwnableID.Bytes()...)
	Bytes = append(Bytes, orderID.TakerOwnableID.Bytes()...)
	Bytes = append(Bytes, rateIDBytes...)
	Bytes = append(Bytes, creationIDBytes...)
	Bytes = append(Bytes, orderID.MakerID.Bytes()...)
	Bytes = append(Bytes, orderID.HashID.Bytes()...)

	return Bytes
}
func (orderID orderID) String() string {
	var values []string
	values = append(values, orderID.ClassificationID.String())
	values = append(values, orderID.MakerOwnableID.String())
	values = append(values, orderID.TakerOwnableID.String())
	values = append(values, orderID.RateID.String())
	values = append(values, orderID.CreationID.String())
	values = append(values, orderID.MakerID.String())
	values = append(values, orderID.HashID.String())

	return strings.Join(values, constants.SecondOrderCompositeIDSeparator)
}
func (orderID orderID) Compare(listable traits.Listable) int {
	return bytes.Compare(orderID.Bytes(), orderIDFromInterface(listable).Bytes())
}
func (orderID orderID) GenerateStoreKeyBytes() []byte {
	return module.StoreKeyPrefix.GenerateStoreKey(orderID.Bytes())
}
func (orderID) RegisterCodec(codec *codec.Codec) {
	codecUtilities.RegisterModuleConcrete(codec, orderID{})
}
func (orderID orderID) IsPartial() bool {
	return len(orderID.HashID.Bytes()) == 0
}
func (orderID orderID) Equals(key helpers.Key) bool {
	return orderID.Compare(orderIDFromInterface(key)) == 0
}

func (orderID orderID) getRateIDBytes() ([]byte, error) {
	var Bytes []byte

	if orderID.RateID.String() == "" {
		return Bytes, nil
	}

	exchangeRate, err := sdkTypes.NewDecFromStr(orderID.RateID.String())
	if err != nil {
		return Bytes, err
	}

	Bytes = append(Bytes, uint8(len(strings.Split(exchangeRate.String(), ".")[0])))
	Bytes = append(Bytes, []byte(exchangeRate.String())...)

	return Bytes, err
}

func (orderID orderID) getCreationHeightBytes() ([]byte, error) {
	var Bytes []byte

	if orderID.CreationID.String() == "" {
		return Bytes, nil
	}

	height, err := strconv.ParseInt(orderID.CreationID.String(), 10, 64)
	if err != nil {
		return Bytes, err
	}

	Bytes = append(Bytes, uint8(len(orderID.CreationID.String())))
	Bytes = append(Bytes, []byte(strconv.FormatInt(height, 10))...)

	return Bytes, err
}

func NewOrderID(classificationID ids.ID, makerOwnableID ids.ID, takerOwnableID ids.ID, rateID ids.ID, creationID ids.ID, makerID ids.ID, immutableProperties lists.PropertyList) ids.ID {
	return orderID{
		ClassificationID: classificationID,
		MakerOwnableID:   makerOwnableID,
		TakerOwnableID:   takerOwnableID,
		RateID:           rateID,
		CreationID:       creationID,
		MakerID:          makerID,
		HashID:           baseQualified.Immutables{PropertyList: immutableProperties}.GenerateHashID(),
	}
}
