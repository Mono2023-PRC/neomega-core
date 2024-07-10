package tiled_buffer

import (
	"bytes"
	"fmt"

	"github.com/OmineDev/neomega-core/minecraft/protocol"
	"github.com/OmineDev/neomega-core/minecraft/protocol/block_nbt"
)

// 将类型为 ID 的方块实体的 __tag NBT 数据从 buffer 底层输出流解码
func Decode(ID string, buffer *bytes.Buffer) (block_nbt.BlockNBT, error) {
	reader := protocol.NewReader(buffer, 0, false)
	block, success := block_nbt.NewPool()[ID]
	if !success {
		return nil, fmt.Errorf("Decode: Can not get target block NBT method; ID = %#v", ID)
	}
	block.Marshal(reader)
	return block, nil
}

// 将 block 编码为 __tag NBT 的二进制数据，
// 同时返回该方块实体对应的 ID 名
func Encode(block block_nbt.BlockNBT) (ID string, bytesGet []byte) {
	buffer := bytes.NewBuffer([]byte{})
	writer := protocol.NewWriter(buffer, 0)
	block.Marshal(writer)
	return block.ID(), buffer.Bytes()
}

// 将 block 写入到一个空切片中，
// 然后从该切片重新阅读数据，
// 并返回该方块实体对应的 NBT 表达形式。
// ID 是该方块实体的 ID 名
func WriteAndRead(ID string, block block_nbt.BlockNBT) (map[string]any, error) {
	id, blockBytes := Encode(block)
	if id != ID {
		return nil, fmt.Errorf("WriteAndRead: ID of block NBT is not matched; id = %#v, ID = %#v", id, ID)
	}
	// write
	new, err := Decode(ID, bytes.NewBuffer(blockBytes))
	if err != nil {
		return nil, fmt.Errorf("WriteAndRead: %v", err)
	}
	// read again
	return new.ToNBT(), nil
	// return
}
