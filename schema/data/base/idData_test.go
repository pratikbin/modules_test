// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package base

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/schema/data/constants"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/utilities/string"
)

func Test_IDData(t *testing.T) {
	idValue := baseIDs.NewStringID("ID")
	testIDData := NewIDData(idValue)
	testIDData2 := NewIDData(baseIDs.NewStringID(""))

	require.Equal(t, "ID", testIDData.String())
	require.Equal(t, baseIDs.NewStringID(string.Hash("ID")), testIDData.GenerateHash())
	require.Equal(t, baseIDs.NewStringID(""), testIDData2.GenerateHash())
	require.Equal(t, constants.IDDataID, testIDData.GetType())

	require.Equal(t, true, NewIDData(baseIDs.NewStringID("identity2")).Compare(NewIDData(baseIDs.NewStringID("identity2"))) == 0)

	require.Panics(t, func() {
		require.Equal(t, false, testIDData.Compare(NewStringData("")) == 0)
	})
	require.Equal(t, true, testIDData.Compare(testIDData) == 0)

	require.Equal(t, "", testIDData.ZeroValue().String())
}