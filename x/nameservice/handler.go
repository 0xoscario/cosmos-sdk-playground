package nameservice

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Now that MsgSetName is specified,
// the next step is to define what action(s) needs to be
// taken when this message is received.
// This is the role of the handler.
// NewHandler returns a handler for "nameservice" type messages
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetName:
			return handleMsgSetName(ctx, keeper, msg)
		case MsgBuyName:
			return handleMsgBuyName(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

// NewHandler is essentially a sub-router
// that directs messages coming into this module to the proper handler.

// In this function, check to see if the Msg sender is actually the owner of the
// name (keeper.GetOwner).
// If so, they can set the name by calling the function on the Keeper.
// If not, throw an error and return that to the user.

func handleMsgSetName(ctx sdk.Context, keeper Keeper, msg MsgSetName) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Name)) {
		return sdk.ErrUnauthorized("incorrect Owner").Result()
	}
	keeper.SetName(ctx, msg.Name, msg.Value)
	return sdk.Result{}
}

func handleMsgBuyName(ctx sdk.Context, keeper Keeper, msg MsgBuyName) sdk.Result {
	// Checks if the the bid price is greater than the price paid by the current owner
	if keeper.GetPrice(ctx, msg.Name).isAllGT(msg.Bid) {
		// If not, throw an error
		return sdk.ErrInsufficientCoins("Bid not high enough").Result()
	}
	if keeper.HasOwner(ctx, msg.Name) {
		_, err = keeper.coinKeeper.SendCoins(ctx, msg.Buyer, keeper.GetOwner(ctx, msg.Name), msg.Bid)

		if err != nil {
			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	} else {
		// if so, deduct the bid amount from the sender
		_, _, err := keeper.coinKeeper.SubtractCoins(ctx, msg.Buyer, msg.Bid)
		if err != nil {

			return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
		}
	}
	keeper.SetOwner(ctx, msg.Name, msg.Buyer)
	keeper.SetPrice(ctx, msg.Name, msg.Buyer)
	return sdk.Result()
}
