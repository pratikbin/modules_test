// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package quash

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/constants/ids"
	"github.com/AssetMantle/modules/modules/identities/auxiliaries/verify"
	"github.com/AssetMantle/modules/modules/identities/internal/key"
	"github.com/AssetMantle/modules/modules/metas/auxiliaries/supplement"
	"github.com/AssetMantle/modules/schema/helpers"
	"github.com/AssetMantle/modules/schema/mappables"
	"github.com/AssetMantle/modules/schema/types/base"
)

type transactionKeeper struct {
	mapper              helpers.Mapper
	supplementAuxiliary helpers.Auxiliary
	verifyAuxiliary     helpers.Auxiliary
}

var _ helpers.TransactionKeeper = (*transactionKeeper)(nil)

func (transactionKeeper transactionKeeper) Transact(context sdkTypes.Context, msg sdkTypes.Msg) helpers.TransactionResponse {
	message := messageFromInterface(msg)

	if auxiliaryResponse := transactionKeeper.verifyAuxiliary.GetKeeper().Help(context, verify.NewAuxiliaryRequest(message.From, message.FromID)); !auxiliaryResponse.IsSuccessful() {
		return newTransactionResponse(auxiliaryResponse.GetError())
	}

	identities := transactionKeeper.mapper.NewCollection(context).Fetch(key.FromID(message.IdentityID))

	identity := identities.Get(key.FromID(message.IdentityID))
	if identity == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	metaProperties, Error := supplement.GetMetaPropertiesFromResponse(transactionKeeper.supplementAuxiliary.GetKeeper().Help(context, supplement.NewAuxiliaryRequest(identity.(mappables.Identity).GetExpiry())))
	if Error != nil {
		return newTransactionResponse(Error)
	}

	expiryHeightMetaFact := metaProperties.Get(ids.ExpiryProperty)
	if expiryHeightMetaFact == nil {
		return newTransactionResponse(errors.EntityNotFound)
	}

	expiryHeight, Error := expiryHeightMetaFact.GetData().AsHeight()
	if Error != nil {
		return newTransactionResponse(Error)
	}

	if expiryHeight.Compare(base.NewHeight(context.BlockHeight())) > 0 {
		return newTransactionResponse(errors.NotAuthorized)
	}

	identities.Remove(identity)

	return newTransactionResponse(nil)
}

func (transactionKeeper transactionKeeper) Initialize(mapper helpers.Mapper, _ helpers.Parameters, auxiliaries []interface{}) helpers.Keeper {
	transactionKeeper.mapper = mapper

	for _, auxiliary := range auxiliaries {
		switch value := auxiliary.(type) {
		case helpers.Auxiliary:
			switch value.GetName() {
			case supplement.Auxiliary.GetName():
				transactionKeeper.supplementAuxiliary = value
			case verify.Auxiliary.GetName():
				transactionKeeper.verifyAuxiliary = value
			}
		default:
			panic(errors.UninitializedUsage)
		}
	}

	return transactionKeeper
}

func keeperPrototype() helpers.TransactionKeeper {
	return transactionKeeper{}
}
