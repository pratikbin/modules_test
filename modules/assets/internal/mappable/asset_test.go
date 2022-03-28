/*
 Copyright [2019] - [2021], PERSISTENCE TECHNOLOGIES PTE. LTD. and the persistenceSDK contributors
 SPDX-License-Identifier: Apache-2.0
*/

package mappable

import (
	"testing"

	"github.com/persistenceOne/persistenceSDK/constants/ids"

	"github.com/stretchr/testify/require"

	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/traits/qualified"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewStringData("ImmutableData")))
	mutableProperties := base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewStringData("MutableData")))

	assetID := key.NewAssetID(classificationID, immutableProperties)
	testAsset := NewAsset(assetID, immutableProperties, mutableProperties).(asset)

	require.Equal(t, asset{Document: qualified.Document{ID: assetID, HasImmutables: qualified.HasImmutables{Properties: immutableProperties}, HasMutables: qualified.HasMutables{Properties: mutableProperties}}}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutableProperties, testAsset.GetImmutableProperties())
	require.Equal(t, mutableProperties, testAsset.GetMutableProperties())
	data, _ := base.ReadHeightData("-1")
	require.Equal(t, testAsset.GetBurn(), base.NewProperty(ids.BurnProperty, data))
	require.Equal(t, base.NewProperty(ids.BurnProperty, base.NewStringData("BurnImmutableData")), asset{qualified.Document{ID: assetID, HasImmutables: qualified.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.BurnProperty, base.NewStringData("BurnImmutableData")))}, HasMutables: qualified.HasMutables{Properties: mutableProperties}}}.GetBurn())
	require.Equal(t, base.NewProperty(ids.BurnProperty, base.NewStringData("BurnMutableData")), asset{qualified.Document{ID: assetID, HasImmutables: qualified.HasImmutables{Properties: immutableProperties}, HasMutables: qualified.HasMutables{Properties: base.NewProperties(base.NewProperty(ids.BurnProperty, base.NewStringData("BurnMutableData")))}}}.GetBurn())
	require.Equal(t, base.NewProperty(ids.LockProperty, data), testAsset.GetLock())
	require.Equal(t, base.NewProperty(ids.LockProperty, base.NewStringData("LockImmutableData")), asset{qualified.Document{ID: assetID, HasImmutables: qualified.HasImmutables{Properties: base.NewProperties(base.NewProperty(ids.LockProperty, base.NewStringData("LockImmutableData")))}, HasMutables: qualified.HasMutables{Properties: mutableProperties}}}.GetLock())
	require.Equal(t, base.NewProperty(ids.LockProperty, base.NewStringData("LockMutableData")), asset{qualified.Document{ID: assetID, HasImmutables: qualified.HasImmutables{Properties: immutableProperties}, HasMutables: qualified.HasMutables{Properties: base.NewProperties(base.NewProperty(ids.LockProperty, base.NewStringData("LockMutableData")))}}}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}