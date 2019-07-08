package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()
	bc.AddBlock("添加一个新区块")
	bc.AddBlock("添加第二个新区块")
	for i, block := range bc.blocks {
		fmt.Printf("当前区块高度: %d\n", i)
		fmt.Printf("时间 %d\n", block.Timestamp)
		fmt.Printf("前区块hash值: %x\n", block.PrevHash)
		fmt.Printf("当前区块hash值: %x\n", block.Hash)
		fmt.Printf("区块数据: %s\n", block.Data)
	}

}
