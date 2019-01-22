package cli

import(
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codect"
	"github.com/cosmos/spf13/cobra"

)

// GetCmdResolveName queries information about a name
func GetCmdResolveName(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use: "resolve [name]",
		Short: "resolve name",
		Args: cobra.ExactArgs(1),
		RunE: func (cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCliContext().WithCodec(cdc)
			name := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, name), nil)
			if err != nil {
				fmt.Printf("could not resolve name - %s \n", string(name))
				return nil
			}

			fmt.Println(string(res))
			return nil
		}
	}
}
