package block_actors

import general "github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/general_actors"

// 孢子花
type SporeBlossom struct {
	general.BlockActor
}

// ID ...
func (*SporeBlossom) ID() string {
	return IDSporeBlossom
}