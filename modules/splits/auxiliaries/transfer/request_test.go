// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package transfer

import (
	"testing"

	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/types/base"
)

func Test_Transfer_Request(t *testing.T) {
	fromID := base.NewID("fromID")
	toID := base.NewID("toID")
	ownableID := base.NewID("ownableID")
	splits := sdkTypes.NewDec(10)
	testAuxiliaryRequest := NewAuxiliaryRequest(fromID, toID, ownableID, splits)

	require.Equal(t, auxiliaryRequest{FromID: fromID, ToID: toID, OwnableID: ownableID, Value: splits}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
