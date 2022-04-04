// Copyright [2021] - [2022], AssetMantle Pte. Ltd. and the code contributors
// SPDX-License-Identifier: Apache-2.0

package key

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AssetMantle/modules/constants"
	baseTraits "github.com/AssetMantle/modules/schema/qualified/base"
	"github.com/AssetMantle/modules/schema/types/base"
	metaUtilities "github.com/AssetMantle/modules/utilities/meta"
)

func Test_ClassificationID_Methods(t *testing.T) {
	chainID := base.NewID("chainID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewStringData("MutableData")))

	testClassificationID := NewClassificationID(chainID, immutableProperties, mutableProperties).(classificationID)
	require.NotPanics(t, func() {
		require.Equal(t, classificationID{ChainID: chainID, HashID: base.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String()))}, testClassificationID)
		require.Equal(t, strings.Join([]string{chainID.String(), base.NewID(metaUtilities.Hash(metaUtilities.Hash("ID1"), metaUtilities.Hash("ID2"), baseTraits.HasImmutables{Properties: immutableProperties}.GenerateHashID().String())).String()}, constants.IDSeparator), testClassificationID.String())
		require.Equal(t, false, testClassificationID.Equals(classificationID{ChainID: base.NewID("chainID"), HashID: base.NewID("hashID")}))
		require.Equal(t, false, testClassificationID.Equals(nil))
		require.Equal(t, false, testClassificationID.Compare(base.NewID("id")) == 0)
		require.Equal(t, true, testClassificationID.Equals(testClassificationID))
		require.Equal(t, false, testClassificationID.IsPartial())
		require.Equal(t, true, classificationID{ChainID: chainID, HashID: base.NewID("")}.IsPartial())
		require.Equal(t, testClassificationID, FromID(testClassificationID))
		require.Equal(t, classificationID{ChainID: base.NewID(""), HashID: base.NewID("")}, FromID(base.NewID("tempID")))
		require.Equal(t, testClassificationID, readClassificationID(testClassificationID.String()))
	})

}
