package types

const (
	// ModuleName defines the module name
	ModuleName = "dex"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_dex"

	// Version defines the current version the IBC module supports
	Version = "dex-1"

	// PortID is the default port id that module binds to
	PortID = "dex"
)

var (
	// PortKey defines the key to store the port ID in store
	PortKey = KeyPrefix("dex-port-")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	VestingKey      = "Vesting/value/"
	VestingCountKey = "Vesting/count/"
)

const (
	SentVestingKey      = "SentVesting/value/"
	SentVestingCountKey = "SentVesting/count/"
)

const (
	TimedoutVestingKey      = "TimedoutVesting/value/"
	TimedoutVestingCountKey = "TimedoutVesting/count/"
)
