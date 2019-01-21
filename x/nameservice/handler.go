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
		default:
			errMsg := fmt.Sprintf("Unrecognized nameservice Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
// NewHandler is essentially a sub-router
// that directs messages coming into this module to the proper handler.

