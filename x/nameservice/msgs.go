package nameservice

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// MsgSetName defines a SetName message
type MsgSetName struct {
	Name  string
	Value string
	Owner sdk.AccAddress
}

type MsgBuyName struct {
	Name  string
	Bid   sdk.Coins
	Buyer sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetName
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name:  name,
		Value: value,
		Owner: owner,
	}
}

// NewMsgBuyName is a constructor function for MsgBuyName
func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name:  name,
		Bid:   bid,
		Buyer: buyer,
	}
}

/////////// Interface

// set
// type should return the name of the module
func (msg MsgSetName) Route() string {
	return "nameservice"
}

// buy
// Type implements Msg.
func (msg MsgBuyName) Route() string {
	return "nameservice"
}

// set
// Name should return the action
func (msg MsgSetName) Type() string {
	return "set_name"
}

// buy
// Name implements Msg.
func (msg MsgBuyName) Type() string {
	return "buy_name"
}

// set
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

// Buy
func (msg MsgBuyName) Validate() sdk.Error {
	if msg.Buyer.Empty() {
		return sdk.ErrUnknownRequest("Name cannot be empty")
	}
	if !msg.Bid.isPostive() {
		return sdk.ErrInsufficientCoins("Bids must be positive")
	}
}

// GetSignBytes Implements Msg.
// GetSignBytes defines how the Msg gets encoded for signing.
// In most cases this means marshal to sorted JSON.
// The output should not be modified.

func (msg MsgSetName) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required on a Tx
// in order for it to be valid.
// In this case, for example,
// the MsgSetName requires that the Owner signs the transaction
// when trying to reset what the name points to.

// GetSigners Implements Msg.
func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}
