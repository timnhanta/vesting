package keeper

import (
	"vesting/x/dex/types"
)

var _ types.QueryServer = Keeper{}
