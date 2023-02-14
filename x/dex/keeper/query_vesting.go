package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"vesting/x/dex/types"
)

func (k Keeper) VestingAll(goCtx context.Context, req *types.QueryAllVestingRequest) (*types.QueryAllVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var vestings []types.Vesting
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	vestingStore := prefix.NewStore(store, types.KeyPrefix(types.VestingKey))

	pageRes, err := query.Paginate(vestingStore, req.Pagination, func(key []byte, value []byte) error {
		var vesting types.Vesting
		if err := k.cdc.Unmarshal(value, &vesting); err != nil {
			return err
		}

		vestings = append(vestings, vesting)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllVestingResponse{Vesting: vestings, Pagination: pageRes}, nil
}

func (k Keeper) Vesting(goCtx context.Context, req *types.QueryGetVestingRequest) (*types.QueryGetVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	vesting, found := k.GetVesting(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetVestingResponse{Vesting: vesting}, nil
}
