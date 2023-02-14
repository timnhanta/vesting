package types

import (
	"fmt"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		PortId:          PortID,
		VestingList:     []Vesting{},
		SentVestingList: []SentVesting{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := host.PortIdentifierValidator(gs.PortId); err != nil {
		return err
	}
	// Check for duplicated ID in vesting
	vestingIdMap := make(map[uint64]bool)
	vestingCount := gs.GetVestingCount()
	for _, elem := range gs.VestingList {
		if _, ok := vestingIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for vesting")
		}
		if elem.Id >= vestingCount {
			return fmt.Errorf("vesting id should be lower or equal than the last id")
		}
		vestingIdMap[elem.Id] = true
	}
	// Check for duplicated ID in sentVesting
	sentVestingIdMap := make(map[uint64]bool)
	sentVestingCount := gs.GetSentVestingCount()
	for _, elem := range gs.SentVestingList {
		if _, ok := sentVestingIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for sentVesting")
		}
		if elem.Id >= sentVestingCount {
			return fmt.Errorf("sentVesting id should be lower or equal than the last id")
		}
		sentVestingIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
