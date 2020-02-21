package main

import (
	"fmt"
)

func main() {
	bc := NewBlockChain()

	for i := 0; i < 10; i++ {

		bc.AddBlock(fmt.Sprintf("添加第%d个区块", i+1))
	}

	for i, block := range bc.blocks {
		fmt.Printf("当前区块高度: %d\n", i)
		fmt.Printf("时间 %d\n", block.Timestamp)
		fmt.Printf("前区块hash值: %x\n", block.PrevHash)
		fmt.Printf("当前区块hash值: %x\n", block.Hash)
		fmt.Printf("区块数据: %s\n", block.Data)
	}

}
