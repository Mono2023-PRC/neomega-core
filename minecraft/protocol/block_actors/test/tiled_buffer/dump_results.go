package tiled_buffer

import "github.com/OmineDev/neomega-core/minecraft/protocol/block_actors"

// ...
type BlockNBTWithBuffer struct {
	ID     string // 该方块实体的 ID 名
	Buffer []byte // 该方块实体的 __tag NBT 二进制切片
}

// 返回一个方块实体的平铺型 __tag NBT 的池，
// 其中包含了一个巨大的列表，其每个元素都记录了
// 方块实体 的 ID 与其对应的
// 平铺型 __tag NBT 的二进制切片。
//
// 这些平铺型的 __tag NBT 将被作为测试用例，
// 用以确定解析库是否可以按预期运行
func NewPool() []BlockNBTWithBuffer {
	return []BlockNBTWithBuffer{
		{
			ID:     block_actors.IDCommandBlock,
			Buffer: []byte{1, 0, 0, 0, 1, 0, 0, 0, 0, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 70, 0, 12, 72, 97, 112, 112, 121, 50, 48, 49, 56, 110, 101, 119, 23, 99, 111, 109, 109, 97, 110, 100, 115, 46, 103, 101, 110, 101, 114, 105, 99, 46, 115, 121, 110, 116, 97, 120, 6, 0, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 1, 0, 254, 177, 11, 0},
		},
		{
			ID:     block_actors.IDCommandBlock,
			Buffer: []byte{1, 0, 0, 0, 1, 1, 0, 0, 0, 26, 115, 99, 111, 114, 101, 98, 111, 97, 114, 100, 32, 111, 98, 106, 101, 99, 116, 105, 118, 101, 115, 32, 108, 105, 115, 116, 70, 0, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 41, 99, 111, 109, 109, 97, 110, 100, 115, 46, 115, 99, 111, 114, 101, 98, 111, 97, 114, 100, 46, 111, 98, 106, 101, 99, 116, 105, 118, 101, 115, 46, 108, 105, 115, 116, 46, 101, 110, 116, 114, 121, 6, 6, 229, 133, 172, 229, 145, 138, 6, 229, 133, 172, 229, 145, 138, 5, 100, 117, 109, 109, 121, 1, 132, 241, 155, 148, 1, 128, 178, 11, 0},
		},
		// 命令方块
		{
			ID:     block_actors.IDNetherReactor,
			Buffer: []byte{1, 0, 0, 1, 128, 3, 1},
		},
		// 下界反应核
		{
			ID:     block_actors.IDSign,
			Buffer: []byte{1, 0, 0, 0, 193, 227, 249, 7, 1, 8, 52, 98, 51, 49, 101, 52, 98, 53, 1, 115, 1, 199, 141, 203, 14, 1, 8, 52, 98, 51, 49, 101, 52, 98, 53, 11, 76, 105, 108, 105, 10, 121, 97, 10, 50, 51, 51, 1},
		},
		// 告示牌
		{
			ID:     block_actors.IDMobSpawner,
			Buffer: []byte{1, 0, 0, 13, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 98, 101, 101, 236, 4, 144, 3, 192, 12, 8, 12, 32, 8, 205, 204, 76, 63, 102, 102, 230, 63, 0, 0, 128, 63, 0, 0},
		},
		// 刷怪笼
		{
			ID:     block_actors.IDSkull,
			Buffer: []byte{1, 0, 0, 4, 0, 0, 52, 66, 0, 0},
		},
		{
			ID:     block_actors.IDSkull,
			Buffer: []byte{1, 0, 0, 5, 0, 0, 52, 66, 1, 250, 10},
		},
		// 头颅
		{
			ID:     block_actors.IDFlowerPot,
			Buffer: []byte{1, 0, 0, 0},
		},
		{
			ID:     block_actors.IDFlowerPot,
			Buffer: []byte{1, 0, 0, 60, 10, 0, 8, 4, 110, 97, 109, 101, 21, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 116, 111, 114, 99, 104, 102, 108, 111, 119, 101, 114, 10, 6, 115, 116, 97, 116, 101, 115, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0},
		},
		// 花盆
		{
			ID:     block_actors.IDEnchantTable,
			Buffer: []byte{1, 0, 0, 0, 0, 135, 234, 1},
		},
		{
			ID:     block_actors.IDEnchantTable,
			Buffer: []byte{1, 0, 0, 0, 255, 174, 231, 191},
		},
		{
			ID:     block_actors.IDEnchantTable,
			Buffer: []byte{1, 4, 84, 101, 115, 116, 0, 0, 69, 174, 35, 192},
		},
		// 附魔台
		{
			ID:     block_actors.IDDayLightDetector,
			Buffer: []byte{1, 0, 0},
		},
		// 阳光探测器
		{
			ID:     block_actors.IDMusic,
			Buffer: []byte{1, 0, 0, 3},
		},
		// 音符盒
		{
			ID:     block_actors.IDComparator,
			Buffer: []byte{1, 0, 0, 30},
		},
		// 比较器
		{
			ID:     block_actors.IDPistonArm,
			Buffer: []byte{0, 0, 0, 0, 0, 128, 63, 0, 0, 128, 63, 2, 2, 1, 0, 0},
		},
		// 活塞
		{
			ID:     block_actors.IDMovingBlock,
			Buffer: []byte{1, 0, 0, 63, 10, 0, 8, 4, 110, 97, 109, 101, 24, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 115, 109, 105, 116, 104, 105, 110, 103, 95, 116, 97, 98, 108, 101, 10, 6, 115, 116, 97, 116, 101, 115, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 54, 10, 0, 8, 4, 110, 97, 109, 101, 15, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 103, 108, 97, 115, 115, 10, 6, 115, 116, 97, 116, 101, 115, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0, 0, 0, 0},
		},
		// 移动的方块
		{
			ID:     block_actors.IDBeacon,
			Buffer: []byte{1, 0, 0, 6, 0},
		},
		// 信标
		{
			ID:     block_actors.IDEndPortal,
			Buffer: []byte{1, 0, 0},
		},
		// 末地折跃门
		{
			ID:     block_actors.IDBed,
			Buffer: []byte{1, 0, 0, 3},
		},
		// 床
		{
			ID:     block_actors.IDBanner,
			Buffer: []byte{1, 0, 0, 15, 0, 4, 3, 99, 114, 101, 10, 3, 99, 98, 111, 12},
		},
		// 旗帜
		{
			ID:     block_actors.IDStructureBlock,
			Buffer: []byte{1, 0, 0, 0, 2, 2, 210, 3, 176, 5, 141, 156, 1, 30, 72, 110, 21, 109, 121, 115, 116, 114, 117, 99, 116, 117, 114, 101, 58, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 200, 66, 0},
		},
		// 结构方块
		/*
			{
				ID:     block_actors.IDChemistryTable,
				Buffer: []byte{}, // 暂缺
			},
			// 化合物创建器
		*/
		{
			ID:     block_actors.IDConduit,
			Buffer: []byte{1, 0, 0, 1, 0},
		},
		// 潮涌核心
		{
			ID:     block_actors.IDJigsawBlock,
			Buffer: []byte{15, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 101, 109, 112, 116, 121, 15, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 101, 109, 112, 116, 121, 15, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 101, 109, 112, 116, 121, 13, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 97, 105, 114, 8, 114, 111, 108, 108, 97, 98, 108, 101, 1, 0, 0},
		},
		// 拼图方块
		{
			ID:     block_actors.IDBell,
			Buffer: []byte{1, 0, 0, 1, 4, 4},
		},
		// 钟
		{
			ID:     block_actors.IDBeehive,
			Buffer: BufferBeehive,
		},
		// 蜂箱
		{
			ID:     block_actors.IDLodestone,
			Buffer: []byte{1, 0, 0, 1, 4},
		},
		{
			ID:     block_actors.IDLodestone,
			Buffer: []byte{1, 0, 0, 0},
		},
		// 磁石
		{
			ID:     block_actors.IDSculkSensor,
			Buffer: []byte{1, 0, 0},
		},
		// 幽匿感测体
		{
			ID:     block_actors.IDSporeBlossom,
			Buffer: []byte{1, 0, 0},
		},
		// 孢子花
		{
			ID:     block_actors.IDSculkCatalyst,
			Buffer: []byte{1, 0, 0},
		},
		// 幽匿催发体
		{
			ID:     block_actors.IDSculkShrieker,
			Buffer: []byte{1, 0, 0},
		},
		// 幽匿感测体
		{
			ID:     block_actors.IDHangingSign,
			Buffer: []byte{1, 0, 0, 1, 193, 227, 249, 7, 1, 8, 52, 98, 51, 49, 101, 52, 98, 53, 12, 233, 173, 148, 230, 179, 149, 228, 185, 166, 239, 188, 129, 0, 255, 255, 255, 15, 1, 8, 52, 98, 51, 49, 101, 52, 98, 53, 6, 228, 189, 160, 229, 165, 189, 0},
		},
		// 悬挂式告示牌
		{
			ID:     block_actors.IDChiseledBookshelf,
			Buffer: []byte{1, 0, 0},
		},
		// 雕纹书架
		{
			ID:     block_actors.IDBrushableBlock,
			Buffer: []byte{0, 1, 0, 0},
		},
		// 可疑的方块
		{
			ID:     block_actors.IDDecoratedPot,
			Buffer: []byte{1, 0, 0},
		},
		// 饰纹陶罐
		{
			ID:     block_actors.IDChest,
			Buffer: []byte{0, 1, 0, 0, 1, 1, 240, 246, 2, 128, 247, 2, 0, 0, 0},
		},
		{
			ID:     block_actors.IDChest,
			Buffer: []byte{0, 1, 0, 0, 1, 0, 0, 0, 180, 1, 8, 29, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 104, 97, 105, 110, 95, 99, 111, 109, 109, 97, 110, 100, 95, 98, 108, 111, 99, 107, 2, 0, 0, 0, 0, 10, 0, 10, 5, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 29, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 104, 97, 105, 110, 95, 99, 111, 109, 109, 97, 110, 100, 95, 98, 108, 111, 99, 107, 10, 6, 115, 116, 97, 116, 101, 115, 1, 15, 99, 111, 110, 100, 105, 116, 105, 111, 110, 97, 108, 95, 98, 105, 116, 0, 3, 16, 102, 97, 99, 105, 110, 103, 95, 100, 105, 114, 101, 99, 116, 105, 111, 110, 0, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0, 16, 21, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 104, 101, 114, 114, 121, 95, 115, 105, 103, 110, 16, 0, 0, 0, 0, 10, 0, 0, 0},
		},
		{
			ID:     block_actors.IDChest,
			Buffer: BufferChest,
		},
		// 箱子
		{
			ID:     block_actors.IDFurnace,
			Buffer: []byte{1, 0, 0, 0, 0, 0, 0, 115, 0, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 97, 109, 111, 110, 100, 95, 115, 119, 111, 114, 100, 1, 0, 0, 0, 6, 2, 17, 0, 3, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 0, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 1, 21, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 104, 101, 114, 114, 121, 95, 115, 105, 103, 110, 16, 0, 0, 0, 0, 10, 0, 0},
		},
		// 熔炉
		{
			ID:     block_actors.IDBrewingStand,
			Buffer: []byte{1, 0, 0, 40, 38, 0, 108, 0, 18, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 114, 101, 100, 115, 116, 111, 110, 101, 10, 0, 0, 0, 0, 10, 0, 0, 1, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 4, 22, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 98, 108, 97, 122, 101, 95, 112, 111, 119, 100, 101, 114, 63, 0, 0, 0, 0, 10, 0, 0},
		},
		// 酿造台
		{
			ID:     block_actors.IDDispenser,
			Buffer: []byte{0, 1, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 248, 1, 0, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 97, 109, 111, 110, 100, 95, 115, 119, 111, 114, 100, 1, 0, 0, 0, 6, 2, 17, 0, 3, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 6, 68, 97, 109, 97, 103, 101, 0, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 0, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 4, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 5, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 6, 21, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 104, 101, 114, 114, 121, 95, 115, 105, 103, 110, 16, 0, 0, 0, 0, 10, 0, 0, 7, 18, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 114, 101, 100, 115, 116, 111, 110, 101, 21, 0, 0, 0, 0, 10, 0, 0, 0},
		},
		// 发射器
		{
			ID:     block_actors.IDDropper,
			Buffer: []byte{0, 1, 0, 0, 147, 3, 0, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 97, 109, 111, 110, 100, 95, 115, 119, 111, 114, 100, 1, 0, 0, 0, 6, 2, 17, 0, 3, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 6, 68, 97, 109, 97, 103, 101, 0, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 0, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 2, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 3, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 114, 111, 112, 112, 101, 114, 64, 0, 0, 0, 0, 10, 0, 10, 5, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 114, 111, 112, 112, 101, 114, 10, 6, 115, 116, 97, 116, 101, 115, 3, 16, 102, 97, 99, 105, 110, 103, 95, 100, 105, 114, 101, 99, 116, 105, 111, 110, 6, 1, 13, 116, 114, 105, 103, 103, 101, 114, 101, 100, 95, 98, 105, 116, 0, 0, 2, 3, 118, 97, 108, 3, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0, 4, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 5, 18, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 114, 101, 100, 115, 116, 111, 110, 101, 21, 0, 0, 0, 0, 10, 0, 0, 7, 22, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 98, 108, 97, 122, 101, 95, 112, 111, 119, 100, 101, 114, 63, 0, 0, 0, 0, 10, 0, 0, 8, 21, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 104, 101, 114, 114, 121, 95, 115, 105, 103, 110, 16, 0, 0, 0, 0, 10, 0, 0, 0},
		},
		// 投掷器
		{
			ID:     block_actors.IDHopper,
			Buffer: []byte{1, 0, 0, 216, 2, 0, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 1, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 114, 111, 112, 112, 101, 114, 64, 0, 0, 0, 0, 10, 0, 10, 5, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 114, 111, 112, 112, 101, 114, 10, 6, 115, 116, 97, 116, 101, 115, 3, 16, 102, 97, 99, 105, 110, 103, 95, 100, 105, 114, 101, 99, 116, 105, 111, 110, 6, 1, 13, 116, 114, 105, 103, 103, 101, 114, 101, 100, 95, 98, 105, 116, 0, 0, 2, 3, 118, 97, 108, 3, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0, 2, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 97, 109, 111, 110, 100, 95, 115, 119, 111, 114, 100, 1, 0, 0, 0, 6, 2, 17, 0, 3, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 6, 68, 97, 109, 97, 103, 101, 0, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 0, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 3, 22, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 98, 108, 97, 122, 101, 95, 112, 111, 119, 100, 101, 114, 63, 0, 0, 0, 0, 10, 0, 0, 4, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 0, 8},
		},
		// 漏斗
		{
			ID:     block_actors.IDCauldron,
			Buffer: []byte{1, 0, 0, 0, 12, 0, 0},
		},
		// 炼药锅
		{
			ID:     block_actors.IDItemFrame,
			Buffer: []byte{1, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 1, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 97, 109, 111, 110, 100, 95, 115, 119, 111, 114, 100, 1, 0, 0, 0, 6, 2, 17, 0, 3, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 6, 68, 97, 109, 97, 103, 101, 0, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 0, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 0, 0, 52, 66, 0, 0, 128, 63},
		},
		{
			ID:     block_actors.IDItemFrame,
			Buffer: []byte{1, 0, 0, 1, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 109, 111, 100, 95, 111, 114, 101, 1, 0, 0, 0, 0, 10, 0, 10, 5, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 109, 111, 100, 95, 111, 114, 101, 10, 6, 115, 116, 97, 116, 101, 115, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 10, 8, 109, 111, 100, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 109, 111, 100, 95, 111, 114, 101, 10, 6, 115, 116, 97, 116, 101, 115, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0, 0, 0, 0, 0, 0, 0, 128, 63},
		},
		// 物品展示框
		{
			ID:     block_actors.IDEnderChest,
			Buffer: []byte{0, 1, 0, 0, 1, 0, 0, 0, 0, 0},
		},
		// 末影箱
		{
			ID:     block_actors.IDShulkerBox,
			Buffer: []byte{2, 0, 1, 0, 0, 1, 0, 0, 0, 253, 2, 2, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 105, 97, 109, 111, 110, 100, 95, 115, 119, 111, 114, 100, 1, 0, 0, 0, 6, 2, 17, 0, 3, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 6, 68, 97, 109, 97, 103, 101, 0, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 0, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 3, 15, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 102, 114, 97, 109, 101, 64, 0, 0, 0, 0, 10, 0, 0, 8, 22, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 98, 108, 97, 122, 101, 95, 112, 111, 119, 100, 101, 114, 63, 0, 0, 0, 0, 10, 0, 0, 12, 22, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 103, 108, 97, 115, 115, 95, 98, 111, 116, 116, 108, 101, 2, 0, 0, 0, 0, 10, 0, 0, 15, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 114, 111, 112, 112, 101, 114, 64, 0, 0, 0, 0, 10, 0, 10, 5, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 17, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 100, 114, 111, 112, 112, 101, 114, 10, 6, 115, 116, 97, 116, 101, 115, 3, 16, 102, 97, 99, 105, 110, 103, 95, 100, 105, 114, 101, 99, 116, 105, 111, 110, 6, 1, 13, 116, 114, 105, 103, 103, 101, 114, 101, 100, 95, 98, 105, 116, 0, 0, 2, 3, 118, 97, 108, 3, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0, 18, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 111, 116, 105, 111, 110, 1, 6, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 1, 13, 119, 97, 115, 74, 117, 115, 116, 66, 114, 101, 119, 101, 100, 1, 0, 0, 26, 18, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 97, 117, 108, 100, 114, 111, 110, 64, 0, 0, 0, 0, 10, 0, 0, 0},
		},
		// 潜影盒
		{
			ID:     block_actors.IDJukebox,
			Buffer: []byte{1, 0, 0, 1, 28, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 109, 117, 115, 105, 99, 95, 100, 105, 115, 99, 95, 112, 105, 103, 115, 116, 101, 112, 1, 0, 0, 0, 0, 10, 0, 0},
		},
		// 唱片机
		{
			ID:     block_actors.IDLectern,
			Buffer: []byte{1, 0, 0, 1, 2, 8, 22, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 119, 114, 105, 116, 116, 101, 110, 95, 98, 111, 111, 107, 1, 0, 0, 0, 0, 10, 0, 10, 3, 116, 97, 103, 8, 6, 97, 117, 116, 104, 111, 114, 3, 42, 42, 42, 3, 10, 103, 101, 110, 101, 114, 97, 116, 105, 111, 110, 0, 9, 5, 112, 97, 103, 101, 115, 10, 8, 8, 9, 112, 104, 111, 116, 111, 110, 97, 109, 101, 0, 8, 4, 116, 101, 120, 116, 6, 230, 181, 139, 232, 175, 149, 0, 8, 9, 112, 104, 111, 116, 111, 110, 97, 109, 101, 0, 8, 4, 116, 101, 120, 116, 7, 230, 181, 139, 232, 175, 149, 50, 0, 8, 9, 112, 104, 111, 116, 111, 110, 97, 109, 101, 0, 8, 4, 116, 101, 120, 116, 6, 228, 189, 160, 229, 165, 189, 0, 8, 9, 112, 104, 111, 116, 111, 110, 97, 109, 101, 0, 8, 4, 116, 101, 120, 116, 9, 228, 184, 150, 231, 149, 140, 239, 188, 129, 0, 8, 5, 116, 105, 116, 108, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 8, 4, 120, 117, 105, 100, 0, 0, 0},
		},
		// 讲台
		{
			ID:     block_actors.IDBlastFurnace,
			Buffer: []byte{1, 0, 0, 0, 0, 0, 0, 134, 1, 0, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 110, 101, 116, 104, 101, 114, 105, 116, 101, 95, 97, 120, 101, 1, 0, 0, 0, 0, 10, 0, 0, 1, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 108, 97, 110, 107, 115, 64, 0, 0, 0, 0, 10, 0, 10, 5, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 108, 97, 110, 107, 115, 10, 6, 115, 116, 97, 116, 101, 115, 8, 9, 119, 111, 111, 100, 95, 116, 121, 112, 101, 3, 111, 97, 107, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0},
		},
		// 高炉
		{
			ID:     block_actors.IDSmoker,
			Buffer: []byte{1, 0, 0, 0, 0, 0, 0, 134, 1, 0, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 110, 101, 116, 104, 101, 114, 105, 116, 101, 95, 97, 120, 101, 1, 0, 0, 0, 0, 10, 0, 0, 1, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 108, 97, 110, 107, 115, 64, 0, 0, 0, 0, 10, 0, 10, 5, 66, 108, 111, 99, 107, 8, 4, 110, 97, 109, 101, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 112, 108, 97, 110, 107, 115, 10, 6, 115, 116, 97, 116, 101, 115, 8, 9, 119, 111, 111, 100, 95, 116, 121, 112, 101, 3, 111, 97, 107, 0, 2, 3, 118, 97, 108, 0, 0, 3, 7, 118, 101, 114, 115, 105, 111, 110, 192, 168, 160, 17, 0, 0},
		},
		// 烟熏炉
		{
			ID:     block_actors.IDCampfire,
			Buffer: []byte{1, 0, 0, 134, 4, 1, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 109, 117, 116, 116, 111, 110, 1, 0, 0, 0, 0, 10, 0, 0, 236, 3, 1, 16, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 109, 117, 116, 116, 111, 110, 1, 0, 0, 0, 0, 10, 0, 0, 204, 3, 1, 14, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 98, 101, 101, 102, 1, 0, 0, 0, 0, 10, 0, 0, 182, 3, 1, 14, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 98, 101, 101, 102, 1, 0, 0, 0, 0, 10, 0, 0},
		},
		// 营火
		{
			ID:     block_actors.IDBarrel,
			Buffer: []byte{0, 1, 0, 0, 1, 0, 0, 0, 228, 1, 0, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 110, 101, 116, 104, 101, 114, 105, 116, 101, 95, 97, 120, 101, 1, 0, 0, 0, 11, 4, 10, 0, 2, 0, 0, 17, 0, 1, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 6, 68, 97, 109, 97, 103, 101, 0, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 2, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 1, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 111, 111, 107, 101, 100, 95, 109, 117, 116, 116, 111, 110, 3, 0, 0, 0, 0, 10, 0, 0, 2, 21, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 111, 111, 107, 101, 100, 95, 98, 101, 101, 102, 3, 0, 0, 0, 0, 10, 0, 0, 3, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 111, 111, 107, 101, 100, 95, 109, 117, 116, 116, 111, 110, 64, 0, 0, 0, 0, 10, 0, 0, 4, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 99, 111, 111, 107, 101, 100, 95, 109, 117, 116, 116, 111, 110, 64, 0, 0, 0, 0, 10, 0, 0, 0},
		},
		// 木桶
		{
			ID:     block_actors.IDGlowItemFrame,
			Buffer: []byte{1, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 1, 23, 109, 105, 110, 101, 99, 114, 97, 102, 116, 58, 110, 101, 116, 104, 101, 114, 105, 116, 101, 95, 97, 120, 101, 1, 0, 0, 0, 11, 4, 10, 0, 2, 0, 0, 17, 0, 1, 0, 0, 10, 0, 10, 3, 116, 97, 103, 3, 6, 68, 97, 109, 97, 103, 101, 0, 3, 10, 82, 101, 112, 97, 105, 114, 67, 111, 115, 116, 2, 10, 7, 100, 105, 115, 112, 108, 97, 121, 8, 4, 78, 97, 109, 101, 9, 76, 105, 108, 105, 121, 97, 50, 51, 51, 0, 0, 0, 0, 0, 7, 67, 0, 0, 128, 63},
		},
		// 荧光物品展示框
		{
			ID:     block_actors.IDCalibratedSculkSensor,
			Buffer: []byte{1, 0, 0},
		},
		// 校频幽匿感测体
		/*
			{
				ID:     block_actors.IDModBlock,
				Buffer: []byte{}, // 暂缺
			},
			// mod_block
		*/
	}
}
