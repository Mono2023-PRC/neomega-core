package bundle

import (
	"fmt"

	"github.com/OmineDev/neomega-core/i18n"
	"github.com/OmineDev/neomega-core/minecraft/protocol/packet"
	"github.com/OmineDev/neomega-core/neomega"
	"github.com/OmineDev/neomega-core/nodes/defines"

	// "github.com/OmineDev/neomega-core/neomega/modules/block/placer"
	"sync"
	"time"

	"github.com/OmineDev/neomega-core/neomega/modules/bot_action"
	"github.com/OmineDev/neomega-core/neomega/modules/chunk_request"
	"github.com/OmineDev/neomega-core/neomega/modules/info_sender"
	"github.com/OmineDev/neomega-core/neomega/modules/player_interact"
	"github.com/OmineDev/neomega-core/neomega/modules/structure"
)

func init() {
	if false {
		func(omega neomega.MicroOmega) {}(&MicroOmega{})
	}
}

type MicroOmega struct {
	neomega.ReactCore
	neomega.InteractCore
	neomega.InfoSender
	neomega.CmdSender
	neomega.MicroUQHolder
	// neomega.BlockPlacer
	neomega.PlayerInteract
	neomega.StructureRequester
	neomega.LowLevelChunkRequester
	neomega.CommandHelper
	neomega.BotAction
	neomega.BotActionHighLevel
	deferredActions []struct {
		cb   func()
		name string
	}
	mu sync.Mutex
}

func NewMicroOmega(
	interactCore neomega.InteractCore,
	reactCore neomega.UnStartedReactCore,
	microUQHolder neomega.MicroUQHolder,
	cmdSender neomega.CmdSender,
	node defines.Node,
	isAccessPoint bool,
) neomega.UnReadyMicroOmega {
	infoSender := info_sender.NewInfoSender(interactCore, cmdSender, microUQHolder.GetBotBasicInfo())
	playerInteract := player_interact.NewPlayerInteract(reactCore, microUQHolder.GetPlayersInfo(), microUQHolder.GetBotBasicInfo(), cmdSender, infoSender, interactCore)
	// asyncNbtBlockPlacer := placer.NewAsyncNbtBlockPlacer(reactCore, cmdSender, interactCore)
	structureRequester := structure.NewStructureRequester(interactCore, reactCore, microUQHolder)
	chunkRequester := chunk_request.NewChunkRequester(interactCore, reactCore, microUQHolder)
	cmdHelper := bot_action.NewCommandHelper(cmdSender, microUQHolder)
	var botAction neomega.BotAction
	if isAccessPoint {
		botAction = bot_action.NewAccessPointBotActionWithPersistData(microUQHolder, interactCore, reactCore, cmdSender, node)
	} else {
		botAction = bot_action.NewEndPointBotAction(node, microUQHolder, interactCore)
	}

	botActionHighLevel := bot_action.NewBotActionHighLevel(microUQHolder, interactCore, reactCore, cmdSender, cmdHelper, structureRequester, botAction, node)

	omega := &MicroOmega{
		reactCore,
		interactCore,
		infoSender,
		cmdSender,
		microUQHolder,
		// asyncNbtBlockPlacer,
		playerInteract,
		structureRequester,
		chunkRequester,
		cmdHelper,
		botAction,
		botActionHighLevel,
		make([]struct {
			cb   func()
			name string
		}, 0),
		sync.Mutex{},
	}

	if isAccessPoint {
		omega.PostponeActionsAfterChallengePassed("request tick update schedule", func() {
			go func() {
				for {
					clientTick := 0
					if tick, found := omega.GetMicroUQHolder().GetExtendInfo().GetCurrentTick(); found {
						clientTick = int(tick)
					}
					omega.GetGameControl().SendPacket(&packet.TickSync{
						ClientRequestTimestamp: int64(clientTick),
					})
					time.Sleep(time.Second * 5)
				}
			}()
		})
	}

	omega.PostponeActionsAfterChallengePassed("dial tick every 1/20 second", func() {
		go func() {
			startTime := time.Now()
			tickAdd := int64(0)
			for {
				// sleep in some platform (yes, you, windows!) is not very accurate
				tickToAdd := (time.Now().Sub(startTime).Milliseconds() / 50) - tickAdd
				if tickToAdd > 0 {
					tickAdd += tickToAdd
					if tick, found := omega.GetMicroUQHolder().GetExtendInfo().GetCurrentTick(); found {
						omega.GetMicroUQHolder().GetExtendInfo().UpdateFromPacket(&packet.TickSync{
							ClientRequestTimestamp:   0,
							ServerReceptionTimestamp: tick + tickToAdd,
						})
					}
				}
				time.Sleep(time.Second / 20)
			}
		}()
	})
	omega.PostponeActionsAfterChallengePassed("force reset dimension and pos", func() {
		e := &neomega.PosAndDimensionInfo{}
		if bot_action.RefreshPosAndDimensionInfo(e, omega) == nil {
			// fmt.Println(e)
			omega.MicroUQHolder.UpdateFromPacket(&packet.ChangeDimension{
				Dimension: int32(e.Dimension),
				Position:  e.HeadPosPrecise,
			})
		}
	})

	if !isAccessPoint {
		omega.PostponeActionsAfterChallengePassed("check bot command status each 10s", func() {
			go func() {
				for {
					ret := omega.SendWebSocketCmdNeedResponse("list").SetTimeout(time.Second * 5).BlockGetResult()
					if ret == nil {
						panic("for some reason, end point cannot communicate with server, reload")
					} else {
						// fmt.Println(ret)
					}
					time.Sleep(time.Second * 10)
				}
			}()
		})
	}

	reactCore.Start()
	return omega
}

func (o *MicroOmega) GetGameControl() neomega.GameCtrl {
	return o
}

func (o *MicroOmega) GetReactCore() neomega.ReactCore {
	return o
}

func (o *MicroOmega) GetGameListener() neomega.PacketDispatcher {
	return o
}

func (o *MicroOmega) GetPlayerInteract() neomega.PlayerInteract {
	return o
}

func (o *MicroOmega) GetMicroUQHolder() neomega.MicroUQHolder {
	return o
}

func (o *MicroOmega) GetStructureRequester() neomega.StructureRequester {
	return o
}

func (o *MicroOmega) GetLowLevelChunkRequester() neomega.LowLevelChunkRequester {
	return o
}

func (o *MicroOmega) GetBotAction() neomega.BotActionComplex {
	return o
}

func (o *MicroOmega) NotifyChallengePassed() {
	for _, action := range o.deferredActions {
		fmt.Printf(i18n.T(i18n.S_starting_post_challenge_actions), action.name)
		action.cb()
	}
}

func (o *MicroOmega) PostponeActionsAfterChallengePassed(name string, action func()) {
	o.mu.Lock()
	defer o.mu.Unlock()
	o.deferredActions = append(o.deferredActions, struct {
		cb   func()
		name string
	}{action, name})
}
