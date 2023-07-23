package keeper

import (
	"context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/UptickNetwork/uptick/x/cw721/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ types.MsgServer = &Keeper{}

// ConvertNFT ConvertCoin converts native Cosmos nft into CW721 tokens for both
// Cosmos-native and CW721 TokenPair Owners
func (k Keeper) ConvertNFT(
	goCtx context.Context,
	msg *types.MsgConvertNFT,
) (
	*types.MsgConvertNFTResponse, error,
) {

	ctx := sdk.UnwrapSDKContext(goCtx)

	//classId, nftIDs
	contractAddress, tokenIds, err := k.GetContractAddressAndTokenIds(ctx, msg)
	if err != nil {
		return nil, err
	}
	msg.ContractAddress = contractAddress
	msg.TokenIds = tokenIds

	// Error checked during msg validation
	receiver := common.HexToAddress(msg.Receiver)

	id := k.GetTokenPairID(ctx, msg.ContractAddress)
	if len(id) == 0 {
		_, err := k.RegisterNFT(ctx, msg)
		if err != nil {
			return nil, err
		}
	}

	pair, err := k.GetPair(ctx, msg.ClassId)
	if err != nil {
		return nil, err
	}

	// Remove token pair if contract is suicided
	cw721 := common.HexToAddress(pair.Cw721Address)
	acc := k.evmKeeper.GetAccountWithoutBalance(ctx, cw721)

	if acc == nil || !acc.IsContract() {
		k.DeleteTokenPair(ctx, pair)
		k.Logger(ctx).Debug(
			"deleting selfdestructed token pair from state",
			"contract", pair.Cw721Address,
		)
		// NOTE: return nil error to persist the changes from the deletion
		return nil, nil
	}
	return k.convertCosmos2Evm(ctx, pair, msg, receiver) // case 2.2
}

// ConvertCW721 converts CW721 tokens into native Cosmos nft for both
// Cosmos-native and CW721 TokenPair Owners
func (k Keeper) ConvertCW721(
	goCtx context.Context,
	msg *types.MsgConvertCW721,
) (
	*types.MsgConvertCW721Response, error,
) {

	ctx := sdk.UnwrapSDKContext(goCtx)
	//classId, nftId
	classId, nftIds, err := k.GetClassIDAndNFTID(ctx, msg)

	fmt.Printf("xxl 0 ConvertCW721 %v-%v \n", classId, nftIds)
	if err != nil {
		return nil, err
	}
	msg.ClassId = classId
	msg.NftIds = nftIds

	// Error checked during msg validation
	sender := common.HexToAddress(msg.Sender)

	id := k.GetTokenPairID(ctx, msg.ContractAddress)
	fmt.Printf("xxl 1 GetTokenPairID %v \n", id)
	if len(id) == 0 {
		_, err := k.RegisterCW721(ctx, msg)
		if err != nil {
			return nil, err
		}
	}

	pair, err := k.GetPair(ctx, msg.ContractAddress)
	fmt.Printf("xxl 2 GetPair %v \n", pair)
	if err != nil {
		return nil, err
	}

	//// Remove token pair if contract is suicided
	//cw721 := common.HexToAddress(pair.Cw721Address)
	//fmt.Printf("xxl 3 cw721 %v \n", cw721)
	//// acc := k.evmKeeper.GetAccountWithoutBalance(ctx, cw721)
	//
	//bigTokenId := new(big.Int)
	//_, err = fmt.Sscan(msg.TokenIds[0], bigTokenId)
	//if err != nil {
	//	sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s error scanning value", err)
	//	return nil, err
	//}

	owner, err := k.QueryCW721TokenOwner(ctx, pair.Cw721Address, msg.TokenIds[0])
	if err != nil {
		return nil, err
	}
	if owner != sender {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of cw721 token %s", sender, strings.Join(msg.TokenIds, ","))
	}

	//if acc == nil || !acc.IsContract() {
	//	k.DeleteTokenPair(ctx, pair)
	//	k.Logger(ctx).Debug(
	//		"deleting selfdestructed token pair from state",
	//		"contract", pair.Cw721Address,
	//	)
	//	// NOTE: return nil error to persist the changes from the deletion
	//	return nil, nil
	//}

	return k.convertEvm2Cosmos(ctx, pair, msg, sender) //

}

// convertCosmos2Evm handles the nft conversion for a native CW721 token
// pair:
//   - escrow nft on module account
//   - unescrow nft that have been previously escrowed with ConvertCW721 and send to receiver
//   - burn escrowed nft
func (k Keeper) convertCosmos2Evm(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertNFT,
	receiver common.Address,
) (
	*types.MsgConvertNFTResponse, error,
) {
	//
	//var (
	//	bigTokenIds []*big.Int
	//	reqInfo     exported.NFT
	//)
	//
	//cw721 := contracts.CW721UpticksContract.ABI
	//contract := pair.GetCW721Contract()
	//
	//for i, tokenId := range msg.TokenIds {
	//	bigTokenId := new(big.Int)
	//	_, err := fmt.Sscan(tokenId, bigTokenId)
	//	if err != nil {
	//		sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s error scanning value", err)
	//		return nil, err
	//	}
	//	bigTokenIds = append(bigTokenIds, bigTokenId)
	//
	//	reqInfo, err = k.nftKeeper.GetNFT(ctx, msg.ClassId, msg.NftIds[i])
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	transferNft := nftTypes.MsgTransferNFT{
	//		DenomId:   msg.ClassId,
	//		Id:        msg.NftIds[i],
	//		Name:      reqInfo.GetName(),
	//		URI:       reqInfo.GetURI(),
	//		Data:      reqInfo.GetData(),
	//		UriHash:   reqInfo.GetURIHash(),
	//		Sender:    msg.Sender,
	//		Recipient: types.AccModuleAddress.String(),
	//	}
	//	if _, err = k.nftKeeper.TransferNFT(ctx, &transferNft); err != nil {
	//		return nil, err
	//	}
	//
	//	//	does token id exist
	//	owner, err := k.QueryCW721TokenOwner(ctx, common.HexToAddress(msg.ContractAddress), bigTokenIds[i])
	//	if err != nil {
	//		// mint
	//		// mint enhance
	//		_, err = k.CallEVM(
	//			ctx, cw721, types.ModuleAddress, contract, true,
	//			"mintEnhance", receiver, bigTokenIds[i], reqInfo.GetName(), reqInfo.GetURI(), reqInfo.GetData(), reqInfo.GetURIHash())
	//		if err != nil {
	//			// mint normal
	//			_, err = k.CallEVM(
	//				ctx, cw721, receiver, contract, true,
	//				"mint", receiver, bigTokenIds[i])
	//			if err != nil {
	//				return nil, err
	//			}
	//		}
	//	} else if owner == types.ModuleAddress {
	//		// transfer
	//		_, err = k.CallEVM(
	//			ctx, cw721, types.ModuleAddress, contract, true,
	//			"safeTransferFrom", types.ModuleAddress, receiver, bigTokenIds[i])
	//		if err != nil {
	//			return nil, err
	//		}
	//	} else {
	//		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of cw721 token %s", types.ModuleAddress, msg.TokenIds)
	//	}
	//
	//	// Mint tokens and send to receiver
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//}
	//
	//for i, tokenId := range msg.TokenIds {
	//	k.SetNFTPairs(ctx, msg.ContractAddress, tokenId, msg.ClassId, msg.NftIds[i])
	//}
	//
	//ctx.EventManager().EmitEvents(
	//	sdk.Events{
	//		sdk.NewEvent(
	//			types.EventTypeConvertNFT,
	//			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
	//			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
	//			sdk.NewAttribute(types.AttributeKeyNFTClass, msg.ClassId),
	//			sdk.NewAttribute(types.AttributeKeyNFTID, strings.Join(msg.NftIds, ",")),
	//			sdk.NewAttribute(types.AttributeKeyCW721Token, contract.String()),
	//			sdk.NewAttribute(types.AttributeKeyCW721TokenID, strings.Join(msg.TokenIds, ",")),
	//		),
	//	},
	//)

	return &types.MsgConvertNFTResponse{}, nil
}

// convertEvm2Cosmos handles the cw721 conversion for a native cw721 token
// pair:
//   - escrow tokens on module account
//   - mint nft to the receiver: nftId: tokenAddress|tokenID
func (k Keeper) convertEvm2Cosmos(
	ctx sdk.Context,
	pair types.TokenPair,
	msg *types.MsgConvertCW721,
	sender common.Address,
) (
	*types.MsgConvertCW721Response, error,
) {
	//
	//cw721 := contracts.CW721UpticksContract.ABI
	//contract := pair.GetCW721Contract()
	//
	//for i, tokenId := range msg.TokenIds {
	//
	//	bigTokenId := new(big.Int)
	//	_, err := fmt.Sscan(tokenId, bigTokenId)
	//	if err != nil {
	//		sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s error scanning value", err)
	//		return nil, err
	//	}
	//
	//	reqInfo, err := k.QueryNFTEnhance(ctx, contract, bigTokenId)
	//	_, err = k.CallEVM(
	//		ctx, cw721, sender, contract, true,
	//		"safeTransferFrom", sender, types.ModuleAddress, bigTokenId,
	//	)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// query cw721 token
	//	_, err = k.QueryCW721Token(ctx, contract)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	nftId := string(k.GetNFTPairByContractTokenID(ctx, msg.ContractAddress, tokenId))
	//	if nftId == "" {
	//
	//		mintNFT := nftTypes.MsgMintNFT{
	//			DenomId:   msg.ClassId,
	//			Id:        msg.NftIds[i],
	//			Name:      reqInfo.Name,
	//			URI:       reqInfo.Uri,
	//			Data:      reqInfo.Data,
	//			UriHash:   reqInfo.UriHash,
	//			Sender:    types.AccModuleAddress.String(),
	//			Recipient: msg.Receiver,
	//		}
	//
	//		// mint nft
	//		if _, err = k.nftKeeper.MintNFT(ctx, &mintNFT); err != nil {
	//			return nil, err
	//		}
	//	} else {
	//		transferNft := nftTypes.MsgTransferNFT{
	//			DenomId:   msg.ClassId,
	//			Id:        msg.NftIds[i],
	//			Name:      reqInfo.Name,
	//			URI:       reqInfo.Uri,
	//			Data:      reqInfo.Data,
	//			UriHash:   reqInfo.UriHash,
	//			Sender:    types.AccModuleAddress.String(),
	//			Recipient: msg.Receiver,
	//		}
	//		if _, err = k.nftKeeper.TransferNFT(ctx, &transferNft); err != nil {
	//			return nil, err
	//		}
	//	}
	//}
	//
	//// save nft pair
	//for i, tokenId := range msg.TokenIds {
	//	k.SetNFTPairs(ctx, msg.ContractAddress, tokenId, msg.ClassId, msg.NftIds[i])
	//}
	//
	//ctx.EventManager().EmitEvents(
	//	sdk.Events{
	//		sdk.NewEvent(
	//			types.EventTypeConvertCW721,
	//			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
	//			sdk.NewAttribute(types.AttributeKeyReceiver, msg.Receiver),
	//			sdk.NewAttribute(types.AttributeKeyNFTClass, pair.ClassId),
	//			sdk.NewAttribute(types.AttributeKeyNFTID, strings.Join(msg.NftIds, ",")),
	//			sdk.NewAttribute(types.AttributeKeyCW721Token, contract.String()),
	//			sdk.NewAttribute(types.AttributeKeyCW721TokenID, strings.Join(msg.TokenIds, ",")),
	//		),
	//	},
	//)

	return &types.MsgConvertCW721Response{}, nil
}