package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Fact(t *testing.T) {

	stringData := NewStringData("testString")
	decData := NewDecData(sdkTypes.NewDec(12))
	idData := NewIDData(NewID("id"))
	heightData := NewHeightData(NewHeight(123))

	testFact := NewFact(stringData)
	require.Equal(t, fact{Hash: stringData.GenerateHash(), Type: "S", Signatures: signatures{}}, testFact)
	require.Equal(t, stringData.GenerateHash(), testFact.GetHash())
	require.Equal(t, signatures{}, testFact.GetSignatures())
	require.Equal(t, false, testFact.(fact).IsMeta())
	require.Equal(t, "S", testFact.GetType())
	require.Equal(t, "D", NewFact(decData).GetType())
	require.Equal(t, "I", NewFact(idData).GetType())
	require.Equal(t, "H", NewFact(heightData).GetType())

	readFact, Error := ReadFact("S|testString")
	require.Equal(t, testFact, readFact)
	require.Nil(t, Error)

	readFact2, Error := ReadFact("")
	require.Equal(t, nil, readFact2)
	require.Equal(t, errors.IncorrectFormat, Error)
}