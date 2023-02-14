package dex_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "vesting/testutil/keeper"
	"vesting/testutil/nullify"
	"vesting/x/dex"
	"vesting/x/dex/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		VestingList: []types.Vesting{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		VestingCount: 2,
		SentVestingList: []types.SentVesting{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentVestingCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DexKeeper(t)
	dex.InitGenesis(ctx, *k, genesisState)
	got := dex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.VestingList, got.VestingList)
	require.Equal(t, genesisState.VestingCount, got.VestingCount)
	require.ElementsMatch(t, genesisState.SentVestingList, got.SentVestingList)
	require.Equal(t, genesisState.SentVestingCount, got.SentVestingCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
