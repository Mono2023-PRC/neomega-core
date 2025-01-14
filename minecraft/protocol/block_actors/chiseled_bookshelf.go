package block_actors

import general "github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/general_actors"

// 雕纹书架
type ChiseledBookshelf struct {
	general.BlockActor `mapstructure:",squash"`
}

// ID ...
func (*ChiseledBookshelf) ID() string {
	return IDChiseledBookshelf
}
