package keeper

import (
	"github.com/defund-labs/osmosis/v11/x/incentives/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams returns all of the parameters in the incentive module.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets all of the parameters in the incentive module.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}
