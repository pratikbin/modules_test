package mappable

import (
	"github.com/persistenceOne/persistenceSDK/constants/properties"
	"github.com/persistenceOne/persistenceSDK/modules/assets/internal/key"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Asset_Methods(t *testing.T) {
	classificationID := base.NewID("classificationID")
	immutables := base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID("ID1"), base.NewFact(base.NewStringData("ImmutableData")))))
	mutables := base.NewMutables(base.NewProperties(base.NewProperty(base.NewID("ID2"), base.NewFact(base.NewStringData("MutableData")))))

	assetID := key.NewAssetID(classificationID, immutables)
	testAsset := NewAsset(assetID, immutables, mutables).(asset)

	require.Equal(t, asset{ID: assetID, Immutables: immutables, Mutables: mutables}, testAsset)
	require.Equal(t, assetID, testAsset.GetID())
	require.Equal(t, classificationID, testAsset.GetClassificationID())
	require.Equal(t, immutables, testAsset.GetImmutables())
	require.Equal(t, mutables, testAsset.GetMutables())
	data, _ := base.ReadHeightData("")
	require.Equal(t, base.NewProperty(base.NewID(properties.Burn), base.NewFact(data)), testAsset.GetBurn())
	require.Equal(t, base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnImmutableData"))), asset{ID: assetID, Immutables: base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnImmutableData"))))), Mutables: mutables}.GetBurn())
	require.Equal(t, base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnMutableData"))), asset{ID: assetID, Immutables: immutables, Mutables: base.NewMutables(base.NewProperties(base.NewProperty(base.NewID(properties.Burn), base.NewFact(base.NewStringData("BurnMutableData")))))}.GetBurn())
	require.Equal(t, base.NewProperty(base.NewID(properties.Lock), base.NewFact(data)), testAsset.GetLock())
	require.Equal(t, base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockImmutableData"))), asset{ID: assetID, Immutables: base.NewImmutables(base.NewProperties(base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockImmutableData"))))), Mutables: mutables}.GetLock())
	require.Equal(t, base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockMutableData"))), asset{ID: assetID, Immutables: immutables, Mutables: base.NewMutables(base.NewProperties(base.NewProperty(base.NewID(properties.Lock), base.NewFact(base.NewStringData("LockMutableData")))))}.GetLock())
	require.Equal(t, assetID, testAsset.GetKey())

}