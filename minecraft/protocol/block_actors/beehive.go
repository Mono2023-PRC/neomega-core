package block_actors

import (
	"github.com/OmineDev/neomega-core/minecraft/protocol"
	"github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/fields"
	general "github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/general_actors"
	"github.com/OmineDev/neomega-core/utils/slices_wrapper"
)

// 蜂箱
type Beehive struct {
	general.BlockActor
	Occupants       []fields.BeehiveOccupants `nbt:"Occupants"`       // TAG_List[TAG_Compound] (9[10])
	ShouldSpawnBees byte                      `nbt:"ShouldSpawnBees"` // TAG_Byte(1) = 0
}

// ID ...
func (*Beehive) ID() string {
	return IDBeehive
}

func (b *Beehive) Marshal(io protocol.IO) {
	protocol.Single(io, &b.BlockActor)
	protocol.SliceVarint16Length(io, &b.Occupants)
	io.Uint8(&b.ShouldSpawnBees)
}

func (b *Beehive) ToNBT() map[string]any {
	var temp map[string]any
	if len(b.Occupants) > 0 {
		new := make([]any, len(b.Occupants))
		for key, value := range b.Occupants {
			new[key] = value.ToNBT()
		}
		temp = map[string]any{
			"Occupants": new,
		}
	}
	return slices_wrapper.MergeMaps(
		b.BlockActor.ToNBT(),
		map[string]any{
			"ShouldSpawnBees": b.ShouldSpawnBees,
		},
		temp,
	)
}

func (b *Beehive) FromNBT(x map[string]any) {
	b.BlockActor.FromNBT(x)
	b.ShouldSpawnBees = x["ShouldSpawnBees"].(byte)

	if occupants, has := x["Occupants"].([]any); has {
		b.Occupants = make([]fields.BeehiveOccupants, len(occupants))
		for key, value := range occupants {
			new := fields.BeehiveOccupants{}
			new.FromNBT(value.(map[string]any))
			b.Occupants[key] = new
		}
	}
}