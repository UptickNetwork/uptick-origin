package keeper

import (
	"fmt"
	nftkeeper "github.com/UptickNetwork/uptick/x/collection/keeper"
	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"github.com/UptickNetwork/uptick/x/erc721/types"
	ibcnfttransferkeeper "github.com/bianjieai/nft-transfer/keeper"

	cw721keep "github.com/UptickNetwork/uptick/x/cw721/keeper"
)

// Keeper of this module maintains collections of erc721.
type Keeper struct {
	storeKey   storetypes.StoreKey
	cdc        codec.BinaryCodec
	paramstore paramtypes.Subspace

	accountKeeper types.AccountKeeper
	nftKeeper     nftkeeper.Keeper
	evmKeeper     types.EVMKeeper
	ics4Wrapper   porttypes.ICS4Wrapper
	ibcKeeper     ibcnfttransferkeeper.Keeper
	cw721Keep     cw721keep.Keeper
}

// NewKeeper creates new instances of the erc721 Keeper
func NewKeeper(storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	nk nftkeeper.Keeper,
	ek types.EVMKeeper,
	ik ibcnfttransferkeeper.Keeper,
) Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		storeKey:      storeKey,
		cdc:           cdc,
		paramstore:    ps,
		accountKeeper: ak,
		nftKeeper:     nk,
		evmKeeper:     ek,
		ibcKeeper:     ik,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetICS4Wrapper sets the ICS4 wrapper to the keeper.
// It panics if already set
func (k *Keeper) SetICS4Wrapper(ics4Wrapper porttypes.ICS4Wrapper) {
	if k.ics4Wrapper != nil {
		panic("ICS4 wrapper already set")
	}

	k.ics4Wrapper = ics4Wrapper
}

// SetCw721Keeper sets the ICS4 wrapper to the keeper.
// It panics if already set
func (k *Keeper) SetCw721Keeper(cw721keeper cw721keep.Keeper) {

	k.cw721Keep = cw721keeper
}
