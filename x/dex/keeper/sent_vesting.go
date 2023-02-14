package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vesting/x/dex/types"
)

// GetSentVestingCount get the total number of sentVesting
func (k Keeper) GetSentVestingCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentVestingCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSentVestingCount set the total number of sentVesting
func (k Keeper) SetSentVestingCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentVestingCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSentVesting appends a sentVesting in the store with a new id and update the count
func (k Keeper) AppendSentVesting(
	ctx sdk.Context,
	sentVesting types.SentVesting,
) uint64 {
	// Create the sentVesting
	count := k.GetSentVestingCount(ctx)

	// Set the ID of the appended value
	sentVesting.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentVestingKey))
	appendedValue := k.cdc.MustMarshal(&sentVesting)
	store.Set(GetSentVestingIDBytes(sentVesting.Id), appendedValue)

	// Update sentVesting count
	k.SetSentVestingCount(ctx, count+1)

	return count
}

// SetSentVesting set a specific sentVesting in the store
func (k Keeper) SetSentVesting(ctx sdk.Context, sentVesting types.SentVesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentVestingKey))
	b := k.cdc.MustMarshal(&sentVesting)
	store.Set(GetSentVestingIDBytes(sentVesting.Id), b)
}

// GetSentVesting returns a sentVesting from its id
func (k Keeper) GetSentVesting(ctx sdk.Context, id uint64) (val types.SentVesting, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentVestingKey))
	b := store.Get(GetSentVestingIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSentVesting removes a sentVesting from the store
func (k Keeper) RemoveSentVesting(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentVestingKey))
	store.Delete(GetSentVestingIDBytes(id))
}

// GetAllSentVesting returns all sentVesting
func (k Keeper) GetAllSentVesting(ctx sdk.Context) (list []types.SentVesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentVestingKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SentVesting
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSentVestingIDBytes returns the byte representation of the ID
func GetSentVestingIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSentVestingIDFromBytes returns ID in uint64 format from a byte array
func GetSentVestingIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
