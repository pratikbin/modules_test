package base

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/helpers"
	"github.com/tendermint/tendermint/libs/kv"
	"testing"
)

func TestMapper(t *testing.T) {

	context, Mapper := setupTest(t)
	testMapper := Mapper.(mapper)

	testMapper.GetKVStoreKey()
	testMapper.NewCollection(context)
	testMapper.Create(context, NewMappable("test1", "value1"))
	testMapper.Create(context, NewMappable("test2", "value2"))
	testMapper.Read(context, NewKey("test1"))
	testMapper.Update(context, NewMappable("test1", "value1"))
	testMapper.Delete(context, NewKey("test2"))
	testMapper.Iterate(context, NewKey("test1"), func(mappable helpers.Mappable) bool { return false })
	testMapper.StoreDecoder(codec.New(), kv.Pair{
		Key: append([]byte{0x11}, []byte("test1")...), Value: testMapper.codec.MustMarshalBinaryBare(NewMappable("test1", "value1"))}, kv.Pair{
		Key: append([]byte{0x11}, []byte("test1")...), Value: testMapper.codec.MustMarshalBinaryBare(NewMappable("test1", "value1"))})

}
