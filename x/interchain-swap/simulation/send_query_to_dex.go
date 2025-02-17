package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/shawayush/inter-swap/x/interchain-swap/keeper"
	"github.com/shawayush/inter-swap/x/interchain-swap/types"
)

func SimulateMsgSendQueryAllBalances(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.QuerySpotPriceStruct{
			Creater: simAccount.Address.String(),
		}

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SendQueryAllBalances simulation not implemented"), nil, nil
	}
}
