package packet

import (
	"github.com/OmineDev/neomega-core/minecraft/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	InputFlagAscend = 1 << iota
	InputFlagDescend
	InputFlagNorthJump
	InputFlagJumpDown
	InputFlagSprintDown
	InputFlagChangeHeight
	InputFlagJumping
	InputFlagAutoJumpingInWater
	InputFlagSneaking
	InputFlagSneakDown
	InputFlagUp
	InputFlagDown
	InputFlagLeft
	InputFlagRight
	InputFlagUpLeft
	InputFlagUpRight
	InputFlagWantUp
	InputFlagWantDown
	InputFlagWantDownSlow
	InputFlagWantUpSlow
	InputFlagSprinting
	InputFlagAscendBlock
	InputFlagDescendBlock
	InputFlagSneakToggleDown
	InputFlagPersistSneak
	InputFlagStartSprinting
	InputFlagStopSprinting
	InputFlagStartSneaking
	InputFlagStopSneaking
	InputFlagStartSwimming
	InputFlagStopSwimming
	InputFlagStartJumping
	InputFlagStartGliding
	InputFlagStopGliding
	InputFlagPerformItemInteraction
	InputFlagPerformBlockActions
	InputFlagPerformItemStackRequest
	InputFlagHandledTeleport
	InputFlagEmoting
	InputFlagMissedSwing
	InputFlagStartCrawling
	InputFlagStopCrawling
)

const (
	InputModeMouse = iota + 1
	InputModeTouch
	InputModeGamePad
	InputModeMotionController
)

const (
	PlayModeNormal = iota
	PlayModeTeaser
	PlayModeScreen
	PlayModeViewer
	PlayModeReality
	PlayModePlacement
	PlayModeLivingRoom
	PlayModeExitLevel
	PlayModeExitLevelLivingRoom
	PlayModeNumModes
)

const (
	InteractionModelTouch = iota
	InteractionModelCrosshair
	InteractionModelClassic
)

// PlayerAuthInput is sent by the client to allow for server authoritative movement. It is used to synchronise
// the player input with the position server-side.
// The client sends this packet when the ServerAuthoritativeMovementMode field in the StartGame packet is set
// to true, instead of the MovePlayer packet. The client will send this packet once every tick.
type PlayerAuthInput struct {
	// Pitch and Yaw hold the rotation that the player reports it has.
	Pitch, Yaw float32
	// Position holds the position that the player reports it has.
	Position mgl32.Vec3
	// MoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of X/Z
	// values which are created using the WASD/controller stick state.
	MoveVector mgl32.Vec2
	// HeadYaw is the horizontal rotation of the head that the player reports it has.
	HeadYaw float32
	// InputData is a combination of bit flags that together specify the way the player moved last tick. It
	// is a combination of the flags above.
	InputData uint64
	// InputMode specifies the way that the client inputs data to the screen. It is one of the constants that
	// may be found above.
	InputMode uint32
	// PlayMode specifies the way that the player is playing. The values it holds, which are rather random,
	// may be found above.
	PlayMode uint32
	// InteractionModel is a constant representing the interaction model the player is using. It is one of the
	// constants that may be found above.
	// Netease
	InteractionModel uint32
	// GazeDirection is the direction in which the player is gazing, when the PlayMode is PlayModeReality: In
	// other words, when the player is playing in virtual reality.
	GazeDirection mgl32.Vec3
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
	// Delta was the delta between the old and the new position. There isn't any practical use for this field
	// as it can be calculated by the server itself.
	Delta mgl32.Vec3
	// ItemInteractionData is the transaction data if the InputData includes an item interaction.
	ItemInteractionData protocol.UseItemTransactionData
	// ItemStackRequest is sent by the client to change an item in their inventory.
	ItemStackRequest protocol.ItemStackRequest
	// BlockActions is a slice of block actions that the client has interacted with.
	BlockActions []protocol.PlayerBlockAction
	// AnalogueMoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of X/Z
	// values which are created using an analogue input.
	AnalogueMoveVector mgl32.Vec2

	/*
		网易专有字段，
		指代玩家的俯仰角 (xBot)。

		该字段看起来与 Pitch 字段是完全相同的，
		目前尚不清楚网易重复这些字段的用途
	*/
	PitchRepeat float32 // Netease
	/*
		网易专有字段，
		指代玩家的偏航角 (yBot)。

		该字段看起来与 Yaw 字段是完全相同的，
		目前尚不清楚网易重复这些字段的用途
	*/
	YawRepeat float32 // Netease
	/*
		网易专有字段，
		可能用于描述玩家是否可以飞行(或正在飞行)。

		正常客户端似乎总是为此提交 false，
		但对于外挂来说，此字段和 CheatOnGround
		同时提交 true 可以在飞行时避免拉回
	*/
	CheatCouldFly bool // Netease
	/*
		网易专有字段，
		可能用于描述玩家是否正在地面上。

		正常客户端似乎总是为此提交 false，
		但对于外挂来说，提交 true 可以在飞行时避免拉回，
		同时，也将成功避免落地伤害
	*/
	CheatOnGround bool // Netease

	// Netease
	Unknown1 bool
}

// ID ...
func (pk *PlayerAuthInput) ID() uint32 {
	return IDPlayerAuthInput
}

func (pk *PlayerAuthInput) Marshal(io protocol.IO) {
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.MoveVector)
	io.Float32(&pk.HeadYaw)
	io.Varuint64(&pk.InputData)
	io.Varuint32(&pk.InputMode)
	io.Varuint32(&pk.PlayMode)
	io.Varuint32(&pk.InteractionModel) // Netease
	if pk.PlayMode == PlayModeReality {
		io.Vec3(&pk.GazeDirection)
	}
	io.Varuint64(&pk.Tick)
	io.Vec3(&pk.Delta)
	io.Bool(&pk.CheatCouldFly) // Netease

	if pk.InputData&InputFlagPerformItemInteraction != 0 {
		io.PlayerInventoryAction(&pk.ItemInteractionData)
	}

	if pk.InputData&InputFlagPerformItemStackRequest != 0 {
		protocol.Single(io, &pk.ItemStackRequest)
	}

	if pk.InputData&InputFlagPerformBlockActions != 0 {
		protocol.SliceVarint32Length(io, &pk.BlockActions)
	}

	io.Vec2(&pk.AnalogueMoveVector)

	// Netease
	io.Bool(&pk.Unknown1)
	io.Float32(&pk.PitchRepeat)
	io.Float32(&pk.YawRepeat)
	io.Bool(&pk.CheatOnGround)
}
