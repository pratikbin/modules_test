// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package conform

import (
	"testing"

	"github.com/stretchr/testify/require"

	baseData "github.com/AssetMantle/modules/schema/data/base"
	baseIDs "github.com/AssetMantle/modules/schema/ids/base"
	"github.com/AssetMantle/modules/schema/lists/base"
	baseProperties "github.com/AssetMantle/modules/schema/properties/base"
)

func Test_Conform_Request(t *testing.T) {
	classificationID := baseIDs.NewID("classificationID")
	mutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID1"), baseData.NewStringData("Data1")))
	immutableProperties := base.NewPropertyList(baseProperties.NewProperty(baseIDs.NewID("ID2"), baseData.NewStringData("Data2")))

	testAuxiliaryRequest := NewAuxiliaryRequest(classificationID, immutableProperties, mutableProperties)

	require.Equal(t, auxiliaryRequest{ClassificationID: classificationID, ImmutableProperties: immutableProperties, MutableProperties: mutableProperties}, testAuxiliaryRequest)
	require.Equal(t, nil, testAuxiliaryRequest.Validate())
	require.Equal(t, testAuxiliaryRequest, auxiliaryRequestFromInterface(testAuxiliaryRequest))
	require.Equal(t, auxiliaryRequest{}, auxiliaryRequestFromInterface(nil))

}
