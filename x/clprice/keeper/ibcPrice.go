package keeper

import (
	"errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/jboetticher/simpleclcosmos/x/clprice/types"
)

// TransmitIbcPricePacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcPricePacket(
	ctx sdk.Context,
	packetData types.IbcPricePacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) error {

	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(channeltypes.ErrChannelNotFound, "port ID (%s) channel ID (%s)", sourcePort, sourceChannel)
	}

	destinationPort := sourceChannelEnd.GetCounterparty().GetPortID()
	destinationChannel := sourceChannelEnd.GetCounterparty().GetChannelID()

	// get the next sequence
	sequence, found := k.channelKeeper.GetNextSequenceSend(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", sourcePort, sourceChannel,
		)
	}

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: "+err.Error())
	}

	packet := channeltypes.NewPacket(
		packetBytes,
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		timeoutHeight,
		timeoutTimestamp,
	)

	if err := k.channelKeeper.SendPacket(ctx, channelCap, packet); err != nil {
		return err
	}

	return nil
}

// OnRecvIbcPricePacket processes packet reception
func (k Keeper) OnRecvIbcPricePacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPricePacketData) (packetAck types.IbcPricePacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic
	id := k.AppendPrice(
		ctx,
		packet.SourcePort+"-"+packet.SourceChannel,
		data.Name,
		data.Price,
		data.Date,
	)
	packetAck.PriceID = strconv.FormatUint(id, 10);

	return packetAck, nil
}

// OnAcknowledgementIbcPricePacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcPricePacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPricePacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcPricePacketAck
		err := packetAck.Unmarshal(dispatchedAck.Result)
		if err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// successful acknowledgement logic
		k.AppendSentPrice(
			ctx,
			packet.SourcePort+"-"+packet.SourceChannel,
			packetAck.PriceID,
			data.Name,
			packet.DestinationPort+"-"+packet.DestinationChannel,
		)

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcPricePacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcPricePacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPricePacketData) error {

	// TODO: packet timeout logic
	// why should we care if they timeout or not? I certainly don't.

	return nil
}
