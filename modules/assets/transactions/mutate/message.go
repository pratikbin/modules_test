package mutate

import (
	"github.com/asaskevich/govalidator"
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/persistenceOne/persistenceSDK/constants"
	"github.com/persistenceOne/persistenceSDK/schema/types"
)

type message struct {
	From       sdkTypes.AccAddress `json:"from" valid:"required~required field fromaddress missing,matches(^commit[a-z0-9]{39}$)~invalid field fromaddress"`
	AssetID    types.ID            `json:"assetid" valid:"required~required field assestid missing"`
	Properties types.Properties    `json:"properties" valid:"required~required field properties missing"`
	Lock       types.Height        `json:"lock" valid:"required~required field lock missing"`
	Burn       types.Height        `json:"burn" valid:"required~required field burn missing"`
}

var _ sdkTypes.Msg = message{}

func (message message) Route() string { return Transaction.GetModuleName() }
func (message message) Type() string  { return Transaction.GetName() }
func (message message) ValidateBasic() error {
	var _, Error = govalidator.ValidateStruct(message)
	if Error != nil {
		return errors.Wrap(constants.IncorrectMessage, Error.Error())
	}
	return nil
}
func (message message) GetSignBytes() []byte {
	return sdkTypes.MustSortJSON(packageCodec.MustMarshalJSON(message))
}
func (message message) GetSigners() []sdkTypes.AccAddress {
	return []sdkTypes.AccAddress{message.From}
}

func messageFromInterface(msg sdkTypes.Msg) message {
	switch value := msg.(type) {
	case message:
		return value
	default:
		return message{}
	}
}

func newMessage(from sdkTypes.AccAddress, assetID types.ID, properties types.Properties, lock types.Height, burn types.Height) sdkTypes.Msg {
	return message{
		From:       from,
		AssetID:    assetID,
		Properties: properties,
		Lock:       lock,
		Burn:       burn,
	}
}
