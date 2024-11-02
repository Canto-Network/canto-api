package queryengine

import (
	"encoding/json"

	math "cosmossdk.io/math"
	inflation "github.com/Canto-Network/Canto/v8/x/inflation/types"
	staking "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// CalculateStakingAPR returns the APR for a all bonded tokens and mint provision for current epoch
func CalculateStakingAPR(pool staking.QueryPoolResponse, mintProvision inflation.QueryEpochMintProvisionResponse) math.LegacyDec {
	//get bonded tokens from pool
	bondedTokens := pool.GetPool().BondedTokens
	//get mint provision amount from epoch (in acanto)
	mintProvisionAmount := mintProvision.GetEpochMintProvision().Amount

	//check if bonded tokens are zero so we don't divide by zero
	if bondedTokens.IsZero() {
		return math.LegacyNewDec(0)
	}

	//calculate apr (mint provision / bonded tokens) * 365 (days) * 100%
	return mintProvisionAmount.Mul(math.LegacyNewDec(36500)).QuoInt(bondedTokens)
}

func GeneralResultToString(results interface{}) string {
	ret, err := json.Marshal(results)
	if err != nil {
		return "GeneralResultToString:" + err.Error()
	}
	return string(ret)
}
