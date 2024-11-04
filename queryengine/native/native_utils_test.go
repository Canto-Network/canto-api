package queryengine

import (
	"reflect"
	"testing"

	math "cosmossdk.io/math"
	inflation "github.com/Canto-Network/Canto/v8/x/inflation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestGetStakingAPR(t *testing.T) {
	type args struct {
		pool          staking.QueryPoolResponse
		mintProvision inflation.QueryEpochMintProvisionResponse
	}
	tests := []struct {
		name string
		args args
		want math.LegacyDec
	}{
		{
			name: "test bonded tokens are zero",
			args: args{
				pool: staking.QueryPoolResponse{
					Pool: staking.Pool{
						BondedTokens: math.ZeroInt(),
					},
				},
				mintProvision: inflation.QueryEpochMintProvisionResponse{
					EpochMintProvision: sdk.NewDecCoin("acanto", math.NewInt(100)),
				},
			},
			want: math.LegacyNewDec(0),
		},
		{
			name: "mint provision is zero",
			args: args{
				pool: staking.QueryPoolResponse{
					Pool: staking.Pool{
						BondedTokens: math.NewInt(100),
					},
				},
				mintProvision: inflation.QueryEpochMintProvisionResponse{
					EpochMintProvision: sdk.NewDecCoin("acanto", math.ZeroInt()),
				},
			},
			want: math.LegacyNewDec(0),
		},
		{
			name: "bonded tokens is less than mint provision",
			args: args{
				pool: staking.QueryPoolResponse{
					Pool: staking.Pool{
						BondedTokens: math.NewInt(100),
					},
				},
				mintProvision: inflation.QueryEpochMintProvisionResponse{
					EpochMintProvision: sdk.NewDecCoin("acanto", math.NewInt(100000000000)),
				},
			},
			want: math.LegacyNewDec(36500000000000),
		},
		{
			name: "mint provision is less than bonded tokens",
			args: args{
				pool: staking.QueryPoolResponse{
					Pool: staking.Pool{
						BondedTokens: math.NewInt(100000000000),
					},
				},
				mintProvision: inflation.QueryEpochMintProvisionResponse{
					EpochMintProvision: sdk.NewDecCoin("acanto", math.NewInt(100)),
				},
			},
			want: math.LegacyMustNewDecFromStr("0.0000365"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculateStakingAPR(tt.args.pool, tt.args.mintProvision); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStakingAPR() = %v, want %v", got, tt.want)
			}
		})
	}
}
