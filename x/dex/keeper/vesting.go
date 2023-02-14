package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"vesting/x/dex/types"
)

// GetVestingCount get the total number of vesting
func (k Keeper) GetVestingCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VestingCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetVestingCount set the total number of vesting
func (k Keeper) SetVestingCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.VestingCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendVesting appends a vesting in the store with a new id and update the count
func (k Keeper) AppendVesting(
	ctx sdk.Context,
	vesting types.Vesting,
) uint64 {
	// Create the vesting
	count := k.GetVestingCount(ctx)

	// Set the ID of the appended value
	vesting.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKey))
	appendedValue := k.cdc.MustMarshal(&vesting)
	store.Set(GetVestingIDBytes(vesting.Id), appendedValue)

	// Update vesting count
	k.SetVestingCount(ctx, count+1)

	return count
}

// SetVesting set a specific vesting in the store
func (k Keeper) SetVesting(ctx sdk.Context, vesting types.Vesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKey))
	b := k.cdc.MustMarshal(&vesting)
	store.Set(GetVestingIDBytes(vesting.Id), b)
}

// GetVesting returns a vesting from its id
func (k Keeper) GetVesting(ctx sdk.Context, id uint64) (val types.Vesting, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKey))
	b := store.Get(GetVestingIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveVesting removes a vesting from the store
func (k Keeper) RemoveVesting(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKey))
	store.Delete(GetVestingIDBytes(id))
}

// GetAllVesting returns all vesting
func (k Keeper) GetAllVesting(ctx sdk.Context) (list []types.Vesting) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VestingKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Vesting
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetVestingIDBytes returns the byte representation of the ID
func GetVestingIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetVestingIDFromBytes returns ID in uint64 format from a byte array
func GetVestingIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
