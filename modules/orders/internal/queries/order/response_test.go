// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package order

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/stretchr/testify/require"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tendermintDB "github.com/tendermint/tm-db"

	"github.com/AssetMantle/modules/constants/errors"
	"github.com/AssetMantle/modules/modules/orders/internal/common"
	"github.com/AssetMantle/modules/modules/orders/internal/mapper"
	"github.com/AssetMantle/modules/schema"
)

func CreateTestInput(t *testing.T) sdkTypes.Context {
	var Codec = codec.New()
	schema.RegisterCodec(Codec)
	sdkTypes.RegisterCodec(Codec)
	codec.RegisterCrypto(Codec)
	codec.RegisterEvidences(Codec)
	vesting.RegisterCodec(Codec)
	Codec.Seal()

	storeKey := sdkTypes.NewKVStoreKey("test")
	paramsStoreKey := sdkTypes.NewKVStoreKey("testParams")
	paramsTransientStoreKeys := sdkTypes.NewTransientStoreKey("testParamsTransient")

	memDB := tendermintDB.NewMemDB()
	commitMultiStore := store.NewCommitMultiStore(memDB)
	commitMultiStore.MountStoreWithDB(storeKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsStoreKey, sdkTypes.StoreTypeIAVL, memDB)
	commitMultiStore.MountStoreWithDB(paramsTransientStoreKeys, sdkTypes.StoreTypeTransient, memDB)
	err := commitMultiStore.LoadLatestVersion()
	require.Nil(t, err)

	context := sdkTypes.NewContext(commitMultiStore, abciTypes.Header{
		ChainID: "test",
	}, false, log.NewNopLogger())

	return context
}

func Test_Order_Response(t *testing.T) {
	context := CreateTestInput(t)
	collection := mapper.Prototype().NewCollection(context)

	testQueryResponse := newQueryResponse(collection, nil)
	testQueryResponseWithError := newQueryResponse(collection, errors.IncorrectFormat)

	require.Equal(t, true, testQueryResponse.IsSuccessful())
	require.Equal(t, false, testQueryResponseWithError.IsSuccessful())
	require.Equal(t, nil, testQueryResponse.GetError())
	require.Equal(t, errors.IncorrectFormat, testQueryResponseWithError.GetError())

	encodedResponse, _ := testQueryResponse.Encode()
	bytes, _ := common.Codec.MarshalJSON(testQueryResponse)
	require.Equal(t, bytes, encodedResponse)

	decodedResponse, _ := queryResponse{}.Decode(bytes)
	require.Equal(t, testQueryResponse, decodedResponse)

	decodedResponse2, _ := queryResponse{}.Decode([]byte{})
	require.Equal(t, nil, decodedResponse2)

	require.Equal(t, queryResponse{}, responsePrototype())
}
