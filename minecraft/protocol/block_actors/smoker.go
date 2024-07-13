package block_actors

import general "github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/general_actors"

// 烟熏炉
type Smoker struct {
	general.FurnaceBlockActor `mapstructure:",squash"`
}

// ID ...
func (*Smoker) ID() string {
	return IDSmoker
}
