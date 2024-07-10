package general

import (
	"github.com/OmineDev/neomega-core/minecraft/protocol"
	"github.com/OmineDev/neomega-core/minecraft/protocol/block_actors/fields"
	"github.com/OmineDev/neomega-core/utils/slices_wrapper"
)

// 描述 物品展示框 和 荧光物品展示框 的通用字段
type ItemFrameBlockActor struct {
	BlockActor
	FrameItem protocol.Optional[fields.FrameItem]
}

func (f *ItemFrameBlockActor) Marshal(r protocol.IO) {
	protocol.Single(r, &f.BlockActor)
	protocol.OptionalMarshaler(r, &f.FrameItem)
}

func (f *ItemFrameBlockActor) ToNBT() map[string]any {
	var temp map[string]any
	if frame, has := f.FrameItem.Value(); has {
		temp = frame.ToNBT()
	}
	return slices_wrapper.MergeMaps(
		f.BlockActor.ToNBT(),
		temp,
	)
}

func (f *ItemFrameBlockActor) FromNBT(x map[string]any) {
	f.BlockActor.FromNBT(x)

	new := fields.FrameItem{}
	if new.CheckExist(x) {
		new.FromNBT(x)
		f.FrameItem = protocol.Optional[fields.FrameItem]{Set: true, Val: new}
	}
}
