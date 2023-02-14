package keeper

import (
	"errors"
	"strconv"

	"vesting/x/dex/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
)

// TransmitIbcVestingPacket transmits the packet over IBC with the specified source port and source channel
func (k Keeper) TransmitIbcVestingPacket(
	ctx sdk.Context,
	packetData types.IbcVestingPacketData,
	sourcePort,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
) (uint64, error) {
	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return 0, sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	packetBytes, err := packetData.GetBytes()
	if err != nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrJSONMarshal, "cannot marshal the packet: %w", err)
	}

	return k.channelKeeper.SendPacket(ctx, channelCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, packetBytes)
}

// OnRecvIbcVestingPacket processes packet reception
func (k Keeper) OnRecvIbcVestingPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcVestingPacketData) (packetAck types.IbcVestingPacketAck, err error) {
	// validate packet data upon receiving
	if err := data.ValidateBasic(); err != nil {
		return packetAck, err
	}

	// TODO: packet reception logic

	id := k.AppendVesting(
		ctx,
		types.Vesting{
			Creator:  packet.SourcePort + "-" + packet.SourceChannel + "-" + data.Creator,
			Start:    data.Start,
			Duration: data.Duration,
			Parts:    data.Parts,
		},
	)

	packetAck.VestingID = strconv.FormatUint(id, 10)

	return packetAck, nil
}

// OnAcknowledgementIbcVestingPacket responds to the the success or failure of a packet
// acknowledgement written on the receiving chain.
func (k Keeper) OnAcknowledgementIbcVestingPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcVestingPacketData, ack channeltypes.Acknowledgement) error {
	switch dispatchedAck := ack.Response.(type) {
	case *channeltypes.Acknowledgement_Error:

		// TODO: failed acknowledgement logic
		_ = dispatchedAck.Error

		return nil
	case *channeltypes.Acknowledgement_Result:
		// Decode the packet acknowledgment
		var packetAck types.IbcVestingPacketAck

		if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
			// The counter-party module doesn't implement the correct acknowledgment format
			return errors.New("cannot unmarshal acknowledgment")
		}

		// TODO: successful acknowledgement logic

		k.AppendSentVesting(
			ctx,
			types.SentVesting{
				Creator:   data.Creator,
				VestingID: packetAck.VestingID,
				Start:     data.Start,
				Duration:  data.Duration,
				Parts:     data.Parts,
				Chain:     packet.DestinationPort + "-" + packet.DestinationChannel,
			},
		)

		return nil
	default:
		// The counter-party module doesn't implement the correct acknowledgment format
		return errors.New("invalid acknowledgment format")
	}
}

// OnTimeoutIbcVestingPacket responds to the case where a packet has not been transmitted because of a timeout
func (k Keeper) OnTimeoutIbcVestingPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcVestingPacketData) error {

	// TODO: packet timeout logic
	k.AppendTimedoutVesting(
		ctx,
		types.TimedoutVesting{
			Creator:  data.Creator,
			Start:    data.Start,
			Duration: data.Duration,
			Parts:    data.Parts,
			Chain:    packet.DestinationPort + "-" + packet.DestinationChannel,
		},
	)

	return nil
}
