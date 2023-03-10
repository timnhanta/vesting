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
				SentVestingList: []types.SentVesting{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				SentVestingCount: 2,
				TimedoutVestingList: []types.TimedoutVesting{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				TimedoutVestingCount: 2,
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
		{
			desc: "duplicated sentVesting",
			genState: &types.GenesisState{
				SentVestingList: []types.SentVesting{
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
			desc: "invalid sentVesting count",
			genState: &types.GenesisState{
				SentVestingList: []types.SentVesting{
					{
						Id: 1,
					},
				},
				SentVestingCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated timedoutVesting",
			genState: &types.GenesisState{
				TimedoutVestingList: []types.TimedoutVesting{
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
			desc: "invalid timedoutVesting count",
			genState: &types.GenesisState{
				TimedoutVestingList: []types.TimedoutVesting{
					{
						Id: 1,
					},
				},
				TimedoutVestingCount: 0,
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
