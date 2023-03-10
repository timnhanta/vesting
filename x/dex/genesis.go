package dex

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vesting/x/dex/keeper"
	"vesting/x/dex/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the vesting
	for _, elem := range genState.VestingList {
		k.SetVesting(ctx, elem)
	}

	// Set vesting count
	k.SetVestingCount(ctx, genState.VestingCount)
	// Set all the sentVesting
	for _, elem := range genState.SentVestingList {
		k.SetSentVesting(ctx, elem)
	}

	// Set sentVesting count
	k.SetSentVestingCount(ctx, genState.SentVestingCount)
	// Set all the timedoutVesting
	for _, elem := range genState.TimedoutVestingList {
		k.SetTimedoutVesting(ctx, elem)
	}

	// Set timedoutVesting count
	k.SetTimedoutVestingCount(ctx, genState.TimedoutVestingCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.VestingList = k.GetAllVesting(ctx)
	genesis.VestingCount = k.GetVestingCount(ctx)
	genesis.SentVestingList = k.GetAllSentVesting(ctx)
	genesis.SentVestingCount = k.GetSentVestingCount(ctx)
	genesis.TimedoutVestingList = k.GetAllTimedoutVesting(ctx)
	genesis.TimedoutVestingCount = k.GetTimedoutVestingCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
