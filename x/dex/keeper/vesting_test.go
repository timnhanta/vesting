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

func createNVesting(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Vesting {
	items := make([]types.Vesting, n)
	for i := range items {
		items[i].Id = keeper.AppendVesting(ctx, items[i])
	}
	return items
}

func TestVestingGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVesting(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetVesting(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestVestingRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVesting(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveVesting(ctx, item.Id)
		_, found := keeper.GetVesting(ctx, item.Id)
		require.False(t, found)
	}
}

func TestVestingGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVesting(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllVesting(ctx)),
	)
}

func TestVestingCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNVesting(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetVestingCount(ctx))
}
