package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "vesting/testutil/keeper"
	"vesting/testutil/nullify"
	"vesting/x/dex/keeper"
	"vesting/x/dex/types"
)

func createNTimedoutVesting(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TimedoutVesting {
	items := make([]types.TimedoutVesting, n)
	for i := range items {
		items[i].Id = keeper.AppendTimedoutVesting(ctx, items[i])
	}
	return items
}

func TestTimedoutVestingGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTimedoutVesting(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTimedoutVesting(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTimedoutVestingRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTimedoutVesting(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTimedoutVesting(ctx, item.Id)
		_, found := keeper.GetTimedoutVesting(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTimedoutVestingGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTimedoutVesting(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTimedoutVesting(ctx)),
	)
}

func TestTimedoutVestingCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNTimedoutVesting(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTimedoutVestingCount(ctx))
}
