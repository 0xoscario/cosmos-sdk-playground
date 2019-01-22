package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ibadsiddiqui/Working-With-Cosmos/x/nameservice"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
 
)

// GetCmdBuyName is the CLI command for sending a BuyName transaction
func GetCmdBuyName(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{ 
		Use: 	"buy-name [name] [amount]"
		Short: 	"bid for existing name or claim new name",
		Args: 	cobra.ExactArgs(2),
		RunE: 	func(cmd *cobra.Command, args []string) error {
					cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

					txBldr := authtxb.NewTxBuilderFromCLI()

					
					if err := cliCtx.EnsureAccountExists(); err != nil {
						return err
					}

					coins, err := sdk.ParseCoins(args[1])
					if err != nil {
						return err
					}

					account, err := cliCtx.GetFromAddress()
					if err != nil {
						return err
					}

					msg := nameservice.NewMsgBuyName(args[0], coins, account)
					err = msg.ValidateBasic()
					if err != nil {
						return err
					}

					cliCtx.PrintResponse = true
					
					return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
				},
	}
}
