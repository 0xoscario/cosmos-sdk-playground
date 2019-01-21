package nameservice

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string
	Value string
	Owner sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

/////////// Interface
// type should return the name of the module
func (msg MsgSetName) Route() string {
	return "nameservice"
}

// Name should return the action
func (msg MsgSetName) Type() string {
	return "set_name"
}

// ValidateBasic Implements Msg.
// ValidateBasic is used to provide some basic stateless checks
// on the validity of the Msg.

func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}
