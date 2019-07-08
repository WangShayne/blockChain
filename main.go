package main

import (
	"crypto/sha256"
	"fmt"
)

//	* 定义区块结构
//	* 前区块哈希
//	* 当前区块哈希
//	* 数据
type Block struct {
	PrevHash []byte //前区块hash
	Hash     []byte //当前区块hash
	Data     []byte //数据
}

//    * 生成哈希
func (block *Block) SetHash() {
	// 获取区块信息
	blockInfo := append(block.PrevHash, block.Data...)
	// 生成hash
	hash := sha256.Sum256(blockInfo)

	block.Hash = hash[:]
}

//   * 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{},
		Data:     []byte(data),
	}

	block.SetHash()

	return &block
}

//  创世区块
func GenesisBlock() *Block {
	return NewBlock("这是一个创世区块", []byte{})
}

//   * 引入区块链
type BlockChain struct {
	blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{
		blocks: []*Block{genesisBlock},
	}
}

//	向链上添加新区块
func (bc *BlockChain) AddBlock(data string) {
	// 获取最后一个区块
	lastBlock := bc.blocks[len(bc.blocks)-1]
	// 获取最后一个区块的hash作为前区块hash
	prevHash := lastBlock.Hash
	// 创建新区块
	block := NewBlock(data, prevHash)
	// 添加新区块
	bc.blocks = append(bc.blocks, block)
}

func main() {
	bc := NewBlockChain()
	bc.AddBlock("添加一个新区块")
	bc.AddBlock("添加第二个新区块")
	for i, block := range bc.blocks {
		fmt.Printf("当前区块高度: %d\n", i)
		fmt.Printf("前区块hash值: %x\n", block.PrevHash)
		fmt.Printf("当前区块hash值: %x\n", block.Hash)
		fmt.Printf("区块数据: %s\n", block.Data)
	}

}
