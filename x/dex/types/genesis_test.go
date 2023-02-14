package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"vesting/x/dex/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
				PortId: types.PortID,
				VestingList: []types.Vesting{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				VestingCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated vesting",
			genState: &types.GenesisState{
				VestingList: []types.Vesting{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid vesting count",
			genState: &types.GenesisState{
				VestingList: []types.Vesting{
					{
						Id: 1,
					},
				},
				VestingCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
