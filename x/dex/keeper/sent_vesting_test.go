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

func createNSentVesting(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SentVesting {
	items := make([]types.SentVesting, n)
	for i := range items {
		items[i].Id = keeper.AppendSentVesting(ctx, items[i])
	}
	return items
}

func TestSentVestingGet(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSentVesting(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSentVesting(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSentVestingRemove(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSentVesting(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSentVesting(ctx, item.Id)
		_, found := keeper.GetSentVesting(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSentVestingGetAll(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSentVesting(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSentVesting(ctx)),
	)
}

func TestSentVestingCount(t *testing.T) {
	keeper, ctx := keepertest.DexKeeper(t)
	items := createNSentVesting(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSentVestingCount(ctx))
}
