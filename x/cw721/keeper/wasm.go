package keeper

import (
	"encoding/json"
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/evmos/ethermint/server/config"
	evmtypes "github.com/evmos/ethermint/x/evm/types"

	"github.com/UptickNetwork/uptick/x/cw721/types"
)

// DeployCW721Contract creates and deploys an CW721 contract on the EVM with the
// erc20 module account as owner.
func (k Keeper) DeployCW721Contract(
	ctx sdk.Context,
	msg *types.MsgConvertNFT,
) (common.Address, error) {

	//class, err := k.nftKeeper.GetDenomInfo(ctx, msg.ClassId)
	//if err != nil {
	//	return common.Address{}, sdkerrors.Wrapf(types.ErrABIPack, "nft class is invalid %s: %s", class.Id, err.Error())
	//}
	//
	//ctorArgs, err := contracts.CW721UpticksContract.ABI.Pack(
	//	"",
	//	class.Name,
	//	class.Symbol,
	//	class.Uri,
	//	class.Data,
	//	class.Description,
	//	class.MintRestricted,
	//	class.Schema,
	//	class.UpdateRestricted,
	//	class.UriHash,
	//)
	//if err != nil {
	//	return common.Address{}, sdkerrors.Wrapf(types.ErrABIPack, "nft class is invalid %s: %s", class.Id, err.Error())
	//}
	//
	//data := make([]byte, len(contracts.CW721UpticksContract.Bin)+len(ctorArgs))
	//copy(data[:len(contracts.CW721UpticksContract.Bin)], contracts.CW721UpticksContract.Bin)
	//copy(data[len(contracts.CW721UpticksContract.Bin):], ctorArgs)
	//
	//nonce, err := k.accountKeeper.GetSequence(ctx, types.ModuleAddress.Bytes())
	//if err != nil {
	//	return common.Address{}, err
	//}
	//
	//contractAddr := crypto.CreateAddress(types.ModuleAddress, nonce)
	//if _, err = k.CallEVMWithData(ctx, types.ModuleAddress, nil, data, true); err != nil {
	//	return common.Address{}, sdkerrors.Wrapf(err, "failed to deploy contract for %s", class.Id)
	//}

	return common.Address{}, nil
}

// QueryCW721 returns the data of a deployed CW721 contract
func (k Keeper) QueryCW721(
	ctx sdk.Context,
	contract common.Address,
) (types.CW721Data, error) {
	//var (
	//	nameRes   types.CW721StringResponse
	//	symbolRes types.CW721StringResponse
	//)

	//cw721 := contracts.CW721UpticksContract.ABI
	//
	//// Name
	//res, err := k.CallEVM(ctx, cw721, types.ModuleAddress, contract, false, "name")
	//if err != nil {
	//	return types.CW721Data{}, err
	//}
	//
	//if err := cw721.UnpackIntoInterface(&nameRes, "name", res.Ret); err != nil {
	//	return types.CW721Data{}, sdkerrors.Wrapf(
	//		types.ErrABIUnpack, "failed to unpack name: %s", err.Error(),
	//	)
	//}
	//
	//// Symbol
	//res, err = k.CallEVM(ctx, cw721, types.ModuleAddress, contract, false, "symbol")
	//if err != nil {
	//	return types.CW721Data{}, err
	//}
	//
	//if err := cw721.UnpackIntoInterface(&symbolRes, "symbol", res.Ret); err != nil {
	//	return types.CW721Data{}, sdkerrors.Wrapf(
	//		types.ErrABIUnpack, "failed to unpack symbol: %s", err.Error(),
	//	)
	//}

	return types.CW721Data{}, nil
}

// QueryClassEnhance returns the data of a deployed CW721 contract
func (k Keeper) QueryClassEnhance(
	ctx sdk.Context,
	contract common.Address,
) (types.ClassEnhance, error) {

	//cw721 := contracts.CW721UpticksContract.ABI
	//
	//// Name
	//res, err := k.CallEVM(ctx, cw721, types.ModuleAddress, contract, false, "getClassEnhanceInfo")
	//if err != nil {
	//	return types.ClassEnhance{}, err
	//}
	//
	//ret, err := cw721.Unpack("getClassEnhanceInfo", res.Ret)
	//if err != nil {
	//	fmt.Printf("QueryClassEnhance resRet %v \n", err)
	//}
	//
	//if len(ret) != 7 {
	//	return types.ClassEnhance{}, nil
	//}
	//
	//return types.NewClassEnhance(
	//	ret[0].(string), ret[1].(string), ret[2].(bool), ret[3].(string),
	//	ret[4].(bool), ret[5].(string), ret[6].(string),
	//), nil

	return types.ClassEnhance{}, nil
}

// QueryNFTEnhance returns the data of a deployed CW721 contract
func (k Keeper) QueryNFTEnhance(
	ctx sdk.Context,
	contract common.Address,
	tokenID *big.Int,
) (types.NFTEnhance, error) {

	//cw721 := contracts.CW721UpticksContract.ABI
	//
	//// Name
	//res, err := k.CallEVM(ctx, cw721, types.ModuleAddress, contract, true, "getNFTEnhanceInfo", tokenID)
	//if err != nil {
	//	return types.NFTEnhance{}, err
	//}
	//
	//ret, err := cw721.Unpack("getNFTEnhanceInfo", res.Ret)
	//if err != nil {
	//	fmt.Printf("QueryNFTEnhance resRet %v \n", err)
	//}
	//
	//if len(ret) != 4 {
	//	return types.NFTEnhance{}, nil
	//}
	//
	//return types.NewNFTEnhance(ret[0].(string), ret[1].(string), ret[2].(string), ret[3].(string)), nil

	return types.NFTEnhance{}, nil
}

// QueryCW721Token returns the data of a CW721 token
func (k Keeper) QueryCW721Token(
	ctx sdk.Context,
	contract common.Address,
) (types.CW721TokenData, error) {
	//var (
	//	nameRes   types.CW721TokenStringResponse
	//	symbolRes types.CW721TokenStringResponse
	//	uriRes    types.CW721TokenStringResponse
	//)
	//
	//cw721 := contracts.CW721UpticksContract.ABI
	//
	//// Name
	//res, err := k.CallEVM(ctx, cw721, types.ModuleAddress, contract, false, "name")
	//if err != nil {
	//	return types.CW721TokenData{}, err
	//}
	//
	//if err := cw721.UnpackIntoInterface(&nameRes, "name", res.Ret); err != nil {
	//	return types.CW721TokenData{}, sdkerrors.Wrapf(
	//		types.ErrABIUnpack, "failed to unpack name: %s", err.Error(),
	//	)
	//}
	//
	//// Symbol
	//res, err = k.CallEVM(ctx, cw721, types.ModuleAddress, contract, false, "symbol")
	//if err != nil {
	//	return types.CW721TokenData{}, err
	//}
	//
	//if err := cw721.UnpackIntoInterface(&symbolRes, "symbol", res.Ret); err != nil {
	//	return types.CW721TokenData{}, sdkerrors.Wrapf(
	//		types.ErrABIUnpack, "failed to unpack symbol: %s", err.Error(),
	//	)
	//}
	//
	//if err := cw721.UnpackIntoInterface(&symbolRes, "symbol", res.Ret); err != nil {
	//	return types.CW721TokenData{}, sdkerrors.Wrapf(
	//		types.ErrABIUnpack, "failed to unpack uri: %s", err.Error(),
	//	)
	//}

	return types.CW721TokenData{}, nil
}

// QueryCW721TokenOwner returns the owner of given tokenID
func (k Keeper) QueryCW721TokenOwner(
	ctx sdk.Context,
	contract string,
	tokenID string,
) (common.Address, error) {

	//var ownerRes types.CW721TokenOwnerResponse
	//
	//cw721 := contracts.CW721UpticksContract.ABI
	//
	//// Name
	//res, err := k.CallEVM(ctx, cw721, types.ModuleAddress, contract, false, "ownerOf", tokenID)
	//if err != nil {
	//	return common.Address{}, err
	//}
	//
	//if err := cw721.UnpackIntoInterface(&ownerRes, "ownerOf", res.Ret); err != nil {
	//	return common.Address{}, sdkerrors.Wrapf(
	//		types.ErrABIUnpack, "failed to unpack owner: %s", err.Error(),
	//	)
	//}
	//
	//return ownerRes.Value, nil

	return common.Address{}, nil
}

// CallEVM performs a smart contract method call using given args
func (k Keeper) CallEVM(
	ctx sdk.Context,
	abi abi.ABI,
	from, contract common.Address,
	commit bool,
	method string,
	args ...interface{},
) (*evmtypes.MsgEthereumTxResponse, error) {

	data, err := abi.Pack(method, args...)
	if err != nil {
		return nil, sdkerrors.Wrap(
			types.ErrABIPack,
			sdkerrors.Wrap(err, "failed to create transaction data").Error(),
		)
	}

	resp, err := k.CallEVMWithData(ctx, from, &contract, data, commit)
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "contract call failed: method '%s', contract '%s'", method, contract)
	}
	return resp, nil
}

// CallEVMWithData performs a smart contract method call using contract data
func (k Keeper) CallEVMWithData(
	ctx sdk.Context,
	from common.Address,
	contract *common.Address,
	data []byte,
	commit bool,
) (*evmtypes.MsgEthereumTxResponse, error) {
	nonce, err := k.accountKeeper.GetSequence(ctx, from.Bytes())
	if err != nil {
		return nil, err
	}

	gasCap := config.DefaultGasCap
	if commit {
		args, err := json.Marshal(evmtypes.TransactionArgs{
			From: &from,
			To:   contract,
			Data: (*hexutil.Bytes)(&data),
		})
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "failed to marshal tx args: %s", err.Error())
		}

		gasRes, err := k.evmKeeper.EstimateGas(sdk.WrapSDKContext(ctx), &evmtypes.EthCallRequest{
			Args:   args,
			GasCap: config.DefaultGasCap,
		})
		if err != nil {
			return nil, err
		}
		gasCap = gasRes.Gas
	}

	msg := ethtypes.NewMessage(
		from,
		contract,
		nonce,
		big.NewInt(0), // amount
		gasCap,        // gasLimit
		big.NewInt(0), // gasFeeCap
		big.NewInt(0), // gasTipCap
		big.NewInt(0), // gasPrice
		data,
		ethtypes.AccessList{}, // AccessList
		!commit,               // isFake
	)

	res, err := k.evmKeeper.ApplyMessage(ctx, msg, evmtypes.NewNoOpTracer(), commit)
	if err != nil {
		return nil, err
	}

	if res.Failed() {
		return nil, sdkerrors.Wrap(evmtypes.ErrVMExecution, res.VmError)
	}

	return res, nil
}
