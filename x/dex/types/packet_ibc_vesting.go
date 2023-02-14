package types

// ValidateBasic is used for validating the packet
func (p IbcVestingPacketData) ValidateBasic() error {

	// TODO: Validate the packet data

	return nil
}

// GetBytes is a helper for serialising
func (p IbcVestingPacketData) GetBytes() ([]byte, error) {
	var modulePacket DexPacketData

	modulePacket.Packet = &DexPacketData_IbcVestingPacket{&p}

	return modulePacket.Marshal()
}
