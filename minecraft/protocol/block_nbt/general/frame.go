package general

import (
	"github.com/OmineDev/neomega-core/minecraft/protocol"
)

// 描述 物品展示框 和 荧光物品展示框 的共用字段
type Frame struct {
	Item           protocol.Item `nbt:"Item"`           // TAG_Compound(10)
	ItemDropChance float32       `nbt:"ItemDropChance"` // TAG_Float(6) = 1
	ItemRotation   float32       `nbt:"ItemRotation"`   // TAG_Float(6) = 0
}

// 检查 x 是否存在 Frame 中记录的所有数据
func (f *Frame) CheckExist(x map[string]any) (exist bool) {
	_, exist1 := x["Item"]
	_, exist2 := x["ItemDropChance"]
	_, exist3 := x["ItemRotation"]
	return exist1 && exist2 && exist3
}

func (f *Frame) Marshal(r protocol.IO) {
	f.Item.Marshal(r)
	r.Float32(&f.ItemRotation)
	r.Float32(&f.ItemDropChance)
}

func (f *Frame) ToNBT() map[string]any {
	return map[string]any{
		"Item":           f.Item.ToNBT(),
		"ItemDropChance": f.ItemDropChance,
		"ItemRotation":   f.ItemRotation,
	}
}

func (f *Frame) FromNBT(x map[string]any) {
	f.Item.FromNBT(x["Item"].(map[string]any))
	f.ItemDropChance = x["ItemDropChance"].(float32)
	f.ItemRotation = x["ItemRotation"].(float32)
}
