package minimal_client_entry

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/OmineDev/neomega-core/i18n"
	"github.com/OmineDev/neomega-core/neomega/rental_server_impact/access_helper"
	"github.com/OmineDev/neomega-core/neomega/rental_server_impact/info_collect_utils"
	"github.com/OmineDev/neomega-core/nodes"
	"github.com/OmineDev/neomega-core/nodes/defines"

	"github.com/pterm/pterm"
)

const ENTRY_NAME = "omega_minimal_client"

func Entry(args *access_helper.ImpactOption) {
	fmt.Println(i18n.T(i18n.S_neomega_access_point_starting))

	if err := info_collect_utils.ReadUserInfoAndUpdateImpactOptions(args); err != nil {
		panic(err)
	}

	accessOption := access_helper.DefaultOptions()
	accessOption.ImpactOption = args
	accessOption.MakeBotCreative = true
	accessOption.DisableCommandBlock = false
	accessOption.ReasonWithPrivilegeStuff = true

	ctx := context.Background()
	node := nodes.NewLocalNode(ctx)
	node = nodes.NewGroup("neomega-core", node, false)
	node.SetValue("_ServerID", defines.FromString(args.ServerCode))
	omegaCore, err := access_helper.ImpactServer(context.Background(), node, accessOption)
	if err != nil {
		panic(err)
	}

	go func() {
		i := 0
		for {
			i++
			time.Sleep(time.Second)
			ret, _ := omegaCore.GetGameControl().SendWebSocketCmdNeedResponse("tp @s ~~~").BlockGetResult()
			if ret == nil || ret.SuccessCount == 0 {
				panic(fmt.Errorf("tp @s ~~~ fail, recv: %v", ret))
			}
		}
	}()
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Printf(">")
			line, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if strings.HasPrefix(line, "/") {
				cmd := strings.TrimPrefix(line, "/")
				ret, _ := omegaCore.GetGameControl().SendWebSocketCmdNeedResponse(cmd).SetTimeout(time.Second).BlockGetResult()
				if ret == nil {
					pterm.Error.Println("cmd not responsed")
				} else {
					bs, _ := json.Marshal(ret)
					pterm.Info.Println(string(bs))
				}
			} else {
				fmt.Println("try type /tp @s ~~~")
			}
		}
	}()
	omegaCore.GetBotAction().DropItemFromHotBar(3)
	panic(<-omegaCore.WaitClosed())
}
