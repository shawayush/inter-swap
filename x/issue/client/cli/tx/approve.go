package tx

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/konstellation/konstellation/x/issue/types"
	"github.com/spf13/cobra"
)

// GetTxCmdApprove Sets `amount` as the allowance of `spender` over the caller's tokens.
func GetTxCmdApprove(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve [spender] [amount]",
		Args:  cobra.ExactArgs(2),
		Short: "Sets `amount` as the allowance of `spender` over the caller's tokens.",
		Long:  "Sets `amount` as the allowance of `spender` over the caller's tokens.",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := auth.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			spender, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			coin, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			msg := types.NewMsgApprove(cliCtx.GetFromAddress(), spender, coin)
			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
