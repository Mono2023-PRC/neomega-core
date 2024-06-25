package packet

import (
	"neo-omega-kernel/minecraft/protocol"
)

// Netease packet
type WithdrawFurnaceXp struct {
	// Netease
	Position protocol.BlockPos
}

// ID ...
func (*WithdrawFurnaceXp) ID() uint32 {
	return IDWithdrawFurnaceXp
}

func (pk *WithdrawFurnaceXp) Marshal(io protocol.IO) {
	io.BlockPos(&pk.Position)
}