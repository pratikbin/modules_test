package schema

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/schema/mappables"
	"github.com/persistenceOne/persistenceSDK/schema/mappers"
	"github.com/persistenceOne/persistenceSDK/schema/types"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/persistenceOne/persistenceSDK/schema/utilities"
)

func RegisterCodec(codec *codec.Codec) {
	types.RegisterCodec(codec)
	base.RegisterCodec(codec)
	mappables.RegisterCodec(codec)
	mappers.RegisterCodec(codec)
	utilities.RegisterCodec(codec)
}
