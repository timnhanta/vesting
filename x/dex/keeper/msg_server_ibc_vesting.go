package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	"vesting/x/dex/types"
)

func (k msgServer) SendIbcVesting(goCtx context.Context, msg *types.MsgSendIbcVesting) (*types.MsgSendIbcVestingResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: logic before transmitting the packet

	// Construct the packet
	var packet types.IbcVestingPacketData

	packet.Start = msg.Start
	packet.Duration = msg.Duration
	packet.Parts = msg.Parts

	// Transmit the packet
	_, err := k.TransmitIbcVestingPacket(
		ctx,
		packet,
		msg.Port,
		msg.ChannelID,
		clienttypes.ZeroHeight(),
		msg.TimeoutTimestamp,
	)
	if err != nil {
		return nil, err
	}

	return &types.MsgSendIbcVestingResponse{}, nil
}
