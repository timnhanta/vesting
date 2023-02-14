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

func (k Keeper) TimedoutVestingAll(goCtx context.Context, req *types.QueryAllTimedoutVestingRequest) (*types.QueryAllTimedoutVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var timedoutVestings []types.TimedoutVesting
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	timedoutVestingStore := prefix.NewStore(store, types.KeyPrefix(types.TimedoutVestingKey))

	pageRes, err := query.Paginate(timedoutVestingStore, req.Pagination, func(key []byte, value []byte) error {
		var timedoutVesting types.TimedoutVesting
		if err := k.cdc.Unmarshal(value, &timedoutVesting); err != nil {
			return err
		}

		timedoutVestings = append(timedoutVestings, timedoutVesting)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTimedoutVestingResponse{TimedoutVesting: timedoutVestings, Pagination: pageRes}, nil
}

func (k Keeper) TimedoutVesting(goCtx context.Context, req *types.QueryGetTimedoutVestingRequest) (*types.QueryGetTimedoutVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	timedoutVesting, found := k.GetTimedoutVesting(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTimedoutVestingResponse{TimedoutVesting: timedoutVesting}, nil
}
