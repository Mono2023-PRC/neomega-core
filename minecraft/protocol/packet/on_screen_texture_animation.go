package packet

import (
	"github.com/OmineDev/neomega-core/minecraft/protocol"
)

// OnScreenTextureAnimation is sent by the server to show a certain animation on the screen of the player.
// The packet is used, as an example, for when a raid is triggered and when a raid is defeated.
type OnScreenTextureAnimation struct {
	// AnimationType is the type of the animation to show. The packet provides no further extra data to allow
	// modifying the duration or other properties of the animation.
	// Netease
	AnimationType uint32
}

// ID ...
func (*OnScreenTextureAnimation) ID() uint32 {
	return IDOnScreenTextureAnimation
}

func (pk *OnScreenTextureAnimation) Marshal(io protocol.IO) {
	io.Uint32(&pk.AnimationType) // For Netease
}
