package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vesting/x/dex/types"
)

// GetTimedoutVestingCount get the total number of timedoutVesting
func (k Keeper) GetTimedoutVestingCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutVestingCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTimedoutVestingCount set the total number of timedoutVesting
func (k Keeper) SetTimedoutVestingCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutVestingCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTimedoutVesting appends a timedoutVesting in the store with a new id and update the count
func (k Keeper) AppendTimedoutVesting(
	ctx sdk.Context,
	timedoutVesting types.TimedoutVesting,
) uint64 {
	// Create the timedoutVesting
	count := k.GetTimedoutVestingCount(ctx)

	// Set the ID of the appended value
	timedoutVesting.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutVestingKey))
	appendedValue := k.cdc.MustMarshal(&timedoutVesting)
	store.Set(GetTimedoutVestingIDBytes(timedoutVesting.Id), appendedValue)

	// Update timedoutVesting count
	k.SetTimedoutVestingCount(ctx, count+1)

	return count
}

// SetTimedoutVesting set a specific timedoutVesting in the store
func (k Keeper) SetTimedoutVesting(ctx sdk.Context, timedoutVesting types.TimedoutVesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutVestingKey))
	b := k.cdc.MustMarshal(&timedoutVesting)
	store.Set(GetTimedoutVestingIDBytes(timedoutVesting.Id), b)
}

// GetTimedoutVesting returns a timedoutVesting from its id
func (k Keeper) GetTimedoutVesting(ctx sdk.Context, id uint64) (val types.TimedoutVesting, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutVestingKey))
	b := store.Get(GetTimedoutVestingIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimedoutVesting removes a timedoutVesting from the store
func (k Keeper) RemoveTimedoutVesting(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutVestingKey))
	store.Delete(GetTimedoutVestingIDBytes(id))
}

// GetAllTimedoutVesting returns all timedoutVesting
func (k Keeper) GetAllTimedoutVesting(ctx sdk.Context) (list []types.TimedoutVesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutVestingKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimedoutVesting
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTimedoutVestingIDBytes returns the byte representation of the ID
func GetTimedoutVestingIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTimedoutVestingIDFromBytes returns ID in uint64 format from a byte array
func GetTimedoutVestingIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
