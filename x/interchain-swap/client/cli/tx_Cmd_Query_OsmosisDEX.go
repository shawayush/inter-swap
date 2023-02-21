package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/shawayush/inter-swap/x/interchain-swap/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdQueryOsmosisDEX() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-pool  [src-channel] [pool-id] [quote-asset-denom] [base-asset-denom]  ",
		Short: "Query osmosis dex pools for tokens ",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgSendQueryAllBalances(
				clientCtx.GetFromAddress().String(),
				args[0], // src-channel
				args[1], // pool-id
				args[2], // quote-asset-denom
				args[3], // base-asset-denom
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "query for DEX pool")

	return cmd
}
