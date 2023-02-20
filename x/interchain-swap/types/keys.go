package types

import (
	icqtypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "interchain-swap"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_interchain-swap"

	PortID = "interquery"

	// Version defines the current version the IBC module supports
	Version = icqtypes.Version
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
