package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSendIbcVesting = "send_ibc_vesting"

var _ sdk.Msg = &MsgSendIbcVesting{}

func NewMsgSendIbcVesting(
	creator string,
	port string,
	channelID string,
	timeoutTimestamp uint64,
	start string,
	duration string,
	parts string,
) *MsgSendIbcVesting {
	return &MsgSendIbcVesting{
		Creator:          creator,
		Port:             port,
		ChannelID:        channelID,
		TimeoutTimestamp: timeoutTimestamp,
		Start:            start,
		Duration:         duration,
		Parts:            parts,
	}
}

func (msg *MsgSendIbcVesting) Route() string {
	return RouterKey
}

func (msg *MsgSendIbcVesting) Type() string {
	return TypeMsgSendIbcVesting
}

func (msg *MsgSendIbcVesting) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSendIbcVesting) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSendIbcVesting) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	if msg.Port == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet port")
	}
	if msg.ChannelID == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet channel")
	}
	if msg.TimeoutTimestamp == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "invalid packet timeout")
	}
	return nil
}
