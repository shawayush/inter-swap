package keeper

import (
	"github.com/shawayush/inter-swap/x/interchain-swap/types"
)

var _ types.QueryServer = Keeper{}
