package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/lavanet/lava/testutil/keeper"
	"github.com/lavanet/lava/testutil/nullify"
	"github.com/lavanet/lava/x/servicer/keeper"
	"github.com/lavanet/lava/x/servicer/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNUniquePaymentStorageUserServicer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.UniquePaymentStorageUserServicer {
	items := make([]types.UniquePaymentStorageUserServicer, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetUniquePaymentStorageUserServicer(ctx, items[i])
	}
	return items
}

func TestUniquePaymentStorageUserServicerGet(t *testing.T) {
	keeper, ctx := keepertest.ServicerKeeper(t)
	items := createNUniquePaymentStorageUserServicer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetUniquePaymentStorageUserServicer(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestUniquePaymentStorageUserServicerRemove(t *testing.T) {
	keeper, ctx := keepertest.ServicerKeeper(t)
	items := createNUniquePaymentStorageUserServicer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveUniquePaymentStorageUserServicer(ctx,
			item.Index,
		)
		_, found := keeper.GetUniquePaymentStorageUserServicer(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestUniquePaymentStorageUserServicerGetAll(t *testing.T) {
	keeper, ctx := keepertest.ServicerKeeper(t)
	items := createNUniquePaymentStorageUserServicer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllUniquePaymentStorageUserServicer(ctx)),
	)
}