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

func (k Keeper) SentVestingAll(goCtx context.Context, req *types.QueryAllSentVestingRequest) (*types.QueryAllSentVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sentVestings []types.SentVesting
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	sentVestingStore := prefix.NewStore(store, types.KeyPrefix(types.SentVestingKey))

	pageRes, err := query.Paginate(sentVestingStore, req.Pagination, func(key []byte, value []byte) error {
		var sentVesting types.SentVesting
		if err := k.cdc.Unmarshal(value, &sentVesting); err != nil {
			return err
		}

		sentVestings = append(sentVestings, sentVesting)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSentVestingResponse{SentVesting: sentVestings, Pagination: pageRes}, nil
}

func (k Keeper) SentVesting(goCtx context.Context, req *types.QueryGetSentVestingRequest) (*types.QueryGetSentVestingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	sentVesting, found := k.GetSentVesting(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSentVestingResponse{SentVesting: sentVesting}, nil
}
