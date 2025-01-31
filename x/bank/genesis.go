package bank

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis sets distribution information for genesis.
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) {
	ctx.Logger().With("module", "x/"+ModuleName).Info("Starting from genesis...")
	keeper.SetSendEnabled(ctx, data.SendEnabled)
	keeper.SetMinSend(ctx, data.MinSend)
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return NewGenesisState(keeper.GetSendEnabled(ctx), keeper.GetMinSend(ctx))
}
