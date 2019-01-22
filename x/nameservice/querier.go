package nameservice

import(
	"github.com/cosmos/cosmos-sdk/codec"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// Whois represents a name -> value lookup
type Whois struct {
	Value string			`json:"value"`
	Owner sdk.AccAddress	`json:"value"`
	Price sdk.Coins			`json:"value"`

}

// query endpoints supported by the governance Querier
// resolve: This takes a name and returns the value that is stored by the nameservice. 
// This is similar to a DNS query.

// whois: This takes a name and returns the price, value, and owner of the name. 
// Used for figuring out how much names cost when you want to buy them.

const (
	QueryResolve = "resolve"
	QueryWhois = "whois"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func (ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryResolve:
			return queryResolve(ctx, path[1], req, keeper)
		case whois:
			reutrn queryWhois(ctx, path[1], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("Unknown nameservice query endpoint")
		}
		
	}
}