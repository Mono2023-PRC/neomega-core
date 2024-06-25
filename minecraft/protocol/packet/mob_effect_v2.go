package packet

import (
	"neo-omega-kernel/minecraft/protocol"
)

// Netease packet
type MobEffectV2 struct {
	// Netease
	Items []protocol.MobEffectV2UnknownEntry
}

// ID ...
func (*MobEffectV2) ID() uint32 {
	return IDMobEffectV2
}

func (pk *MobEffectV2) Marshal(io protocol.IO) {
	protocol.SliceVaruint32Length(io, &pk.Items)
}