package main

import "crypto/sha256"

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
