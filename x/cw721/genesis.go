package cw721

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	"github.com/UptickNetwork/uptick/x/cw721/keeper"
	"github.com/UptickNetwork/uptick/x/cw721/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	data types.GenesisState,
) {
	fmt.Printf("xxl cw721 1 InitGenesis%v \n", data.Params)
	k.SetParams(ctx, data.Params)

	// ensure cw721 module account is set on genesis
	if acc := accountKeeper.GetModuleAccount(ctx, types.ModuleName); acc == nil {
		panic("the cw721 module account has not been set")
	}

	fmt.Printf("xxl cw721 2 data.TokenPairs %v \n", data.TokenPairs)

	for _, pair := range data.TokenPairs {
		id := pair.GetID()
		k.SetTokenPair(ctx, pair)
		k.SetClassMap(ctx, pair.ClassId, id)
		k.SetCW721Map(ctx, pair.GetCw721Address(), id)

		fmt.Printf("xxl cw721 3 id %v \n", id)
		fmt.Printf("xxl cw721 4 pair %v \n", pair)
	}
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params:     k.GetParams(ctx),
		TokenPairs: k.GetTokenPairs(ctx),
	}
}
