package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/shawayush/inter-swap/x/interchain-swap/keeper"
	keepertest "github.com/shawayush/inter-swap/x/interchain-swap/keeper"
	"github.com/shawayush/inter-swap/x/interchain-swap/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.MarsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
