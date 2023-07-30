package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

// constants
const (

	// DefaultPrefix prefix
	DefaultPrefix = "uptick"

	// ModuleName module name
	ModuleName = "cw721"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for message routing
	RouterKey = ModuleName
)

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address
var AccModuleAddress sdk.AccAddress

func init() {
	AccModuleAddress = authtypes.NewModuleAddress(ModuleName)
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

// prefix bytes for the EVM persistent store
const (
	prefixTokenPair = iota + 1
	prefixTokenPairByCW721
	prefixTokenPairByClass

	prefixNFTUIDPairByNFTUID
	prefixNFTUIDPairByTokenUID

	prefixWasmCode
)

// KVStore key prefixes
var (
	KeyPrefixTokenPair        = []byte{prefixTokenPair}
	KeyPrefixTokenPairByCW721 = []byte{prefixTokenPairByCW721}
	KeyPrefixTokenPairByClass = []byte{prefixTokenPairByClass}

	KeyPrefixNFTUIDPairByNFTUID   = []byte{prefixNFTUIDPairByNFTUID}
	KeyPrefixNFTUIDPairByTokenUID = []byte{prefixNFTUIDPairByTokenUID}

	KeyPrefixWasmCode = []byte{prefixWasmCode}
)
