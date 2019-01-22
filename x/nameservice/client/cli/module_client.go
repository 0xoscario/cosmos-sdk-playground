package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	nameservicecmd "github.com/ibadsiddiqui/Working-With-Cosmos/x/nameservice/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exoirts all client functionality from this module
type ModuleClient struct {
	storekey string
	cdc *amino.Codec
}

func NewModuleClient(storekey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient(storekey, cdc)
}

//