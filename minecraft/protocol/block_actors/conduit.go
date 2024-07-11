package block_actors

import (
	"github.com/OmineDev/neomega-core/minecraft/protocol"
	general "github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/general_actors"
	"github.com/OmineDev/neomega-core/utils/slices_wrapper"
)

// 潮涌核心
type Conduit struct {
	general.BlockActor
	Active byte  `nbt:"Active"` // TAG_Byte(1) = 0
	Target int64 `nbt:"Target"` // TAG_Long(5) = -1
}

// ID ...
func (*Conduit) ID() string {
	return IDConduit
}

func (c *Conduit) Marshal(io protocol.IO) {
	protocol.Single(io, &c.BlockActor)
	io.Varint64(&c.Target)
	io.Uint8(&c.Active)
}

func (c *Conduit) ToNBT() map[string]any {
	return slices_wrapper.MergeMaps(
		c.BlockActor.ToNBT(),
		map[string]any{
			"Active": c.Active,
			"Target": c.Target,
		},
	)
}

func (c *Conduit) FromNBT(x map[string]any) {
	c.BlockActor.FromNBT(x)
	c.Active = x["Active"].(byte)
	c.Target = x["Target"].(int64)
}