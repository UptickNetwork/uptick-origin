package erc721

import (
	"fmt"
	"github.com/bianjieai/nft-transfer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v5/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v5/modules/core/exported"
	"strings"

	"github.com/UptickNetwork/uptick/ibc"
	"github.com/UptickNetwork/uptick/x/erc721/keeper"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the transfer middleware given
// the claim keeper and the underlying application.
type IBCMiddleware struct {
	*ibc.Module
	keeper keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application
func NewIBCMiddleware(k keeper.Keeper, app porttypes.IBCModule) IBCMiddleware {
	return IBCMiddleware{
		Module: ibc.NewModule(app),
		keeper: k,
	}
}

// OnRecvPacket implements the IBCModule interface.
// If fees are not enabled, this callback will default to the ibc-core packet callback.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {

	fmt.Printf("###xxl 721 IBCMiddleware OnRecvPacket \n")
	ack := im.Module.OnRecvPacket(ctx, packet, relayer)

	// return if the acknowledgement is an error ACK
	if !ack.Success() {
		return ack
	}

	return nil
}

// OnAcknowledgementPacket implements the IBCModule interface
// If fees are not enabled, this callback will default to the ibc-core packet callback.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {

	// decode the data
	var ack channeltypes.Acknowledgement
	if err := types.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal ICS-721 transfer packet acknowledgement: %v", err)
	}

	var data types.NonFungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest,
			"cannot unmarshal ICS-721 transfer packet data: %s", err.Error())
	}

	if err := im.keeper.OnAcknowledgementPacket(ctx, packet, data, ack); err != nil {
		return err
	}

	fmt.Printf("###xxl 721 IBCMiddleware OnAcknowledgementPacket \n")
	if err := im.Module.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer); err != nil {

		fmt.Printf("###xxl 721 IBCMiddleware OnAcknowledgementPacket %v \n", err)
		return err
	}

	return nil
}

// SendPacket implements the ICS4 Wrapper interface
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
) error {
	return nil
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet exported.PacketI,
	ack exported.Acknowledgement,
) error {
	fmt.Printf("###xxl 721 IBCMiddleware WriteAcknowledgement \n")
	return nil
}

func (im IBCMiddleware) GetAppVersion(
	ctx sdk.Context,
	portID,
	channelID string,
) (string, bool) {
	return "", false
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	var data types.NonFungibleTokenPacketData
	if err := types.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-721 transfer packet data: %s", err.Error())
	}
	// refund tokens
	if err := im.keeper.OnTimeoutPacket(ctx, packet, data); err != nil {
		return err
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeTimeout,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(types.AttributeKeySender, data.Sender),
			sdk.NewAttribute(types.AttributeKeyReceiver, data.Receiver),
			sdk.NewAttribute(types.AttributeKeyClassID, data.ClassId),
			sdk.NewAttribute(types.AttributeKeyTokenIDs, strings.Join(data.TokenIds, ",")),
		),
	)

	return nil
}
