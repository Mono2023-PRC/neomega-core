package block_actors

import general "github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/general_actors"

// 末影箱
type EnderChest struct {
	general.ChestBlockActor
}

// ID ...
func (*EnderChest) ID() string {
	return IDEnderChest
}