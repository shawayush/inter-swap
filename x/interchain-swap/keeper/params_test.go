package keeper_test

import (
	"testing"

	testkeeper "github.com/shawayush/inter-swap/x/interchain-swap/keeper"
	"github.com/shawayush/inter-swap/x/interchain-swap/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MarsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
