package bot_action

import (
	"bytes"
	"time"

	"github.com/OmineDev/neomega-core/minecraft/protocol"
	"github.com/OmineDev/neomega-core/neomega"
	"github.com/OmineDev/neomega-core/neomega/chunks/define"
	"github.com/OmineDev/neomega-core/nodes"
	"github.com/OmineDev/neomega-core/nodes/defines"
)

type EndPointBotAction struct {
	*BotActionSimple
	node defines.Node
}

func (a *AccessPointBotActionWithPersistData) ExposeAPI() {
	a.ExposeInventoryContent()
	a.ExposeSelectHotBar()
	a.ExposeUseHotBarItemOnBlock()
	a.ExposeUseHotBarItemOnBlockWithOffset()
	a.ExposeMoveItemFromInventoryToEmptyContainerSlots()
	a.ExposeUseAnvil()
	a.ExposeDropItemFromHotBar()
	a.ExposeMoveItemInsideHotBarOrInventory()
	a.ExposeUseHotBarItem()
}

func NewEndPointBotAction(node defines.Node, uq neomega.MicroUQHolder, ctrl neomega.InteractCore) neomega.BotAction {
	ba := &EndPointBotAction{
		BotActionSimple: NewBotActionSimple(uq, ctrl),
		node:            nodes.NewGroup("bot_action", node, false),
	}
	return ba
}

func (a *AccessPointBotActionWithPersistData) ExposeInventoryContent() {
	a.node.ExposeAPI("get_inventory_content", func(args defines.Values) (result defines.Values, err error) {
		var windowID uint32
		var slot uint8
		if err = (&ArgsChain{resArgs: args}).TakeUint32(&windowID).TakeUint8(&slot).Error(); err != nil {
			return defines.Empty, err
		}
		inventoryContent, found := a.GetInventoryContent(windowID, slot)
		if !found {
			return defines.Empty, nil
		}
		buf := bytes.NewBuffer([]byte{})
		writer := protocol.NewWriter(buf, 0)
		writer.ItemInstance(inventoryContent)
		return defines.FromFrags(buf.Bytes()), nil
	}, true)
}

func (e *EndPointBotAction) GetInventoryContent(windowID uint32, slotID uint8) (*protocol.ItemInstance, bool) {
	args := (&ArgsChain{}).SetUint32(windowID).SetUint8(slotID).Done()
	ret, err := e.node.CallWithResponse("get_inventory_content", args).SetTimeout(time.Second * 30).BlockGetResult()
	if err != nil || ret.IsEmpty() {
		return nil, false
	}
	data, err := ret.ToBytes()
	if err != nil {
		return nil, false
	}
	buf := bytes.NewBuffer(data)
	reader := protocol.NewReader(buf, 0, false)
	instance := protocol.ItemInstance{}
	reader.ItemInstance(&instance)
	return &instance, true
}

func (a *AccessPointBotActionWithPersistData) ExposeUseHotBarItemOnBlock() {
	a.node.ExposeAPI("use_hot_bar_item_on_block", func(args defines.Values) (result defines.Values, err error) {
		var blockPos define.CubePos
		var blockNEMCRuntimeID uint32
		var face int32
		var slot uint8
		if err = (&ArgsChain{resArgs: args}).TakePos(&blockPos).TakeUint32(&blockNEMCRuntimeID).TakeInt32(&face).TakeUint8(&slot).Error(); err != nil {
			return defines.Empty, err
		}
		err = a.UseHotBarItemOnBlock(blockPos, blockNEMCRuntimeID, face, slot)
		return defines.Empty, err
	}, true)
}

func (a *AccessPointBotActionWithPersistData) ExposeUseHotBarItemOnBlockWithOffset() {
	a.node.ExposeAPI("use_hot_bar_item_on_block_with_offset", func(args defines.Values) (result defines.Values, err error) {
		var blockPos define.CubePos
		var blockNEMCRuntimeID uint32
		var face int32
		var slot uint8
		var offsetPos define.CubePos
		if err = (&ArgsChain{resArgs: args}).TakePos(&blockPos).TakePos(&offsetPos).TakeUint32(&blockNEMCRuntimeID).TakeInt32(&face).TakeUint8(&slot).Error(); err != nil {
			return defines.Empty, err
		}
		err = a.UseHotBarItemOnBlockWithBotOffset(blockPos, offsetPos, blockNEMCRuntimeID, face, slot)
		return defines.Empty, err
	}, true)
}

func (e *EndPointBotAction) UseHotBarItemOnBlock(blockPos define.CubePos, blockNEMCRuntimeID uint32, face int32, slot uint8) (err error) {
	args := (&ArgsChain{}).SetPos(blockPos).SetUint32(blockNEMCRuntimeID).SetInt32(face).SetUint8(slot).Done()
	_, err = e.node.CallWithResponse("use_hot_bar_item_on_block", args).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

func (e *EndPointBotAction) UseHotBarItemOnBlockWithBotOffset(blockPos define.CubePos, botOffset define.CubePos, blockNEMCRuntimeID uint32, face int32, slot uint8) (err error) {
	args := (&ArgsChain{}).SetPos(blockPos).SetPos(botOffset).SetUint32(blockNEMCRuntimeID).SetInt32(face).SetUint8(slot).Done()
	_, err = e.node.CallWithResponse("use_hot_bar_item_on_block_with_offset", args).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

func (a *AccessPointBotActionWithPersistData) ExposeSelectHotBar() {
	a.node.ExposeAPI("select_hot_bar", func(args defines.Values) (result defines.Values, err error) {
		var slot uint8
		if err = (&ArgsChain{resArgs: args}).TakeUint8(&slot).Error(); err != nil {
			return defines.Empty, err
		}
		err = a.selectHotBar(slot)
		return defines.Empty, err
	}, true)
}

func (e *EndPointBotAction) SelectHotBar(slotID uint8) error {
	args := (&ArgsChain{}).SetUint8(slotID).Done()
	_, err := e.node.CallWithResponse("select_hot_bar", args).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

func (a *AccessPointBotActionWithPersistData) ExposeMoveItemFromInventoryToEmptyContainerSlots() {
	a.node.ExposeAPI("move_item_from_inventory_slot_to_empty_container_slots", func(args defines.Values) (result defines.Values, err error) {
		var pos define.CubePos
		var blockNEMCRuntimeID uint32
		var blockName string
		var counts uint8
		chain := (&ArgsChain{resArgs: args})
		if err = chain.TakePos(&pos).TakeUint32(&blockNEMCRuntimeID).TakeString(&blockName).TakeUint8(&counts).Error(); err != nil {
			return defines.Empty, err
		}
		switchOperations := map[uint8]uint8{}
		for i := 0; i < int(counts); i++ {
			var k, v uint8
			if err = chain.TakeUint8(&k).TakeUint8(&v).Error(); err != nil {
				return defines.Empty, err
			}
			switchOperations[k] = v
		}
		err = a.MoveItemFromInventoryToEmptyContainerSlots(pos, blockNEMCRuntimeID, blockName, switchOperations)
		return defines.Empty, err
	}, true)
}

func (e *EndPointBotAction) MoveItemFromInventoryToEmptyContainerSlots(pos define.CubePos, blockNemcRtid uint32, blockName string, switchOperations map[uint8]uint8) error {
	args := (&ArgsChain{}).SetPos(pos).SetUint32(blockNemcRtid).SetString(blockName).SetUint8(uint8(len(switchOperations)))
	for k, v := range switchOperations {
		args.SetUint8(k).SetUint8(v)
	}
	_, err := e.node.CallWithResponse("move_item_from_inventory_slot_to_empty_container_slots", args.Done()).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

func (a *AccessPointBotActionWithPersistData) ExposeUseAnvil() {
	a.node.ExposeAPI("use_anvil", func(args defines.Values) (result defines.Values, err error) {
		var pos define.CubePos
		var blockNEMCRuntimeID uint32
		var slot uint8
		var newName string
		if err = (&ArgsChain{resArgs: args}).TakePos(&pos).TakeUint32(&blockNEMCRuntimeID).TakeUint8(&slot).TakeString(&newName).Error(); err != nil {
			return defines.Empty, err
		}
		err = a.UseAnvil(pos, blockNEMCRuntimeID, slot, newName)
		return defines.Empty, err
	}, true)
}

func (e *EndPointBotAction) UseAnvil(pos define.CubePos, blockNemcRtid uint32, slot uint8, newName string) error {
	args := (&ArgsChain{}).SetPos(pos).SetUint32(blockNemcRtid).SetUint8(slot).SetString(newName).Done()
	_, err := e.node.CallWithResponse("use_anvil", args).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

func (a *AccessPointBotActionWithPersistData) ExposeDropItemFromHotBar() {
	a.node.ExposeAPI("drop_item_from_hot_bar", func(args defines.Values) (result defines.Values, err error) {
		var slot uint8
		if err = (&ArgsChain{resArgs: args}).TakeUint8(&slot).Error(); err != nil {
			return defines.Empty, err
		}
		err = a.DropItemFromHotBar(slot)
		return defines.Empty, err
	}, true)
}

func (e *EndPointBotAction) DropItemFromHotBar(slot uint8) error {
	args := (&ArgsChain{}).SetUint8(slot).Done()
	_, err := e.node.CallWithResponse("drop_item_from_hot_bar", args).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

func (a *AccessPointBotActionWithPersistData) ExposeMoveItemInsideHotBarOrInventory() {
	a.node.ExposeAPI("move_item_inside_hotbar_or_inventory", func(args defines.Values) (result defines.Values, err error) {
		var sourceSlot, targetSlot, count uint8
		if err = (&ArgsChain{resArgs: args}).TakeUint8(&sourceSlot).TakeUint8(&targetSlot).TakeUint8(&count).Error(); err != nil {
			return defines.Empty, err
		}
		err = a.MoveItemInsideHotBarOrInventory(sourceSlot, targetSlot, count)
		return defines.Empty, err
	}, true)
}

func (e *EndPointBotAction) MoveItemInsideHotBarOrInventory(sourceSlot, targetSlot, count uint8) error {
	args := (&ArgsChain{}).SetUint8(sourceSlot).SetUint8(targetSlot).SetUint8(count).Done()
	_, err := e.node.CallWithResponse("move_item_inside_hotbar_or_inventory", args).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

func (a *AccessPointBotActionWithPersistData) ExposeUseHotBarItem() {
	a.node.ExposeAPI("use_hotbar_item", func(args defines.Values) (result defines.Values, err error) {
		var slot uint8
		if err = (&ArgsChain{resArgs: args}).TakeUint8(&slot).Error(); err != nil {
			return defines.Empty, err
		}
		err = a.UseHotBarItem(slot)
		return defines.Empty, err
	}, true)
}

func (e *EndPointBotAction) UseHotBarItem(slot uint8) (err error) {
	args := (&ArgsChain{}).SetUint8(slot).Done()
	_, err = e.node.CallWithResponse("use_hotbar_item", args).SetTimeout(time.Second * 30).BlockGetResult()
	return err
}

type ArgsChain struct {
	err     error
	resArgs defines.Values
}

func (c *ArgsChain) Error() error {
	return c.err
}

func (c *ArgsChain) Done() defines.Values {
	return c.resArgs
}

func (c *ArgsChain) SetInt64(x int64) *ArgsChain {
	if c.resArgs == nil {
		c.resArgs = make(defines.Values, 0)
	}
	c.resArgs = c.resArgs.Extend(defines.FromInt64(x))
	return c
}

func (c *ArgsChain) TakeInt64(x *int64) *ArgsChain {
	if c.err == nil {
		*x, c.err = c.resArgs.ToInt64()
		c.resArgs = c.resArgs.ConsumeHead()
	}
	return c
}

func (c *ArgsChain) SetInt32(x int32) *ArgsChain {
	if c.resArgs == nil {
		c.resArgs = make(defines.Values, 0)
	}
	c.resArgs = c.resArgs.Extend(defines.FromInt32(x))
	return c
}

func (c *ArgsChain) TakeInt32(x *int32) *ArgsChain {
	if c.err == nil {
		*x, c.err = c.resArgs.ToInt32()
		c.resArgs = c.resArgs.ConsumeHead()
	}
	return c
}

func (c *ArgsChain) SetUint64(x uint64) *ArgsChain {
	if c.resArgs == nil {
		c.resArgs = make(defines.Values, 0)
	}
	c.resArgs = c.resArgs.Extend(defines.FromUint64(x))
	return c
}

func (c *ArgsChain) TakeUint64(x *int64) *ArgsChain {
	if c.err == nil {
		*x, c.err = c.resArgs.ToInt64()
		c.resArgs = c.resArgs.ConsumeHead()
	}
	return c
}

func (c *ArgsChain) SetUint32(x uint32) *ArgsChain {
	if c.resArgs == nil {
		c.resArgs = make(defines.Values, 0)
	}
	c.resArgs = c.resArgs.Extend(defines.FromUint32(x))
	return c
}

func (c *ArgsChain) TakeUint32(x *uint32) *ArgsChain {
	if c.err == nil {
		*x, c.err = c.resArgs.ToUint32()
		c.resArgs = c.resArgs.ConsumeHead()
	}
	return c
}

func (c *ArgsChain) SetUint8(x byte) *ArgsChain {
	if c.resArgs == nil {
		c.resArgs = make(defines.Values, 0)
	}
	c.resArgs = c.resArgs.Extend(defines.FromByte(x))
	return c
}

func (c *ArgsChain) TakeUint8(x *byte) *ArgsChain {
	if c.err == nil {
		*x, c.err = c.resArgs.ToByte()
		c.resArgs = c.resArgs.ConsumeHead()
	}
	return c
}

func (c *ArgsChain) SetPos(pos define.CubePos) *ArgsChain {
	return c.SetInt64(int64(pos.X())).SetInt64(int64(pos.Y())).SetInt64(int64(pos.Z()))
}

func (c *ArgsChain) TakePos(pos *define.CubePos) *ArgsChain {
	if c.err == nil {
		var x, y, z int64
		c.TakeInt64(&x).TakeInt64(&y).TakeInt64(&z)
		if c.err == nil {
			*pos = define.CubePos{int(x), int(y), int(z)}
		}
	}
	return c
}

func (c *ArgsChain) SetString(x string) *ArgsChain {
	if c.resArgs == nil {
		c.resArgs = make(defines.Values, 0)
	}
	c.resArgs = c.resArgs.Extend(defines.FromString(x))
	return c
}

func (c *ArgsChain) TakeString(x *string) *ArgsChain {
	if c.err == nil {
		*x, c.err = c.resArgs.ToString()
		c.resArgs = c.resArgs.ConsumeHead()
	}
	return c
}

func (c *ArgsChain) SetBytes(x []byte) *ArgsChain {
	if c.resArgs == nil {
		c.resArgs = make(defines.Values, 0)
	}
	c.resArgs = c.resArgs.ExtendFrags(x)
	return c
}

func (c *ArgsChain) TakeBytes(x *[]byte) *ArgsChain {
	if c.err == nil {
		*x, c.err = c.resArgs.ToBytes()
		c.resArgs = c.resArgs.ConsumeHead()
	}
	return c
}
