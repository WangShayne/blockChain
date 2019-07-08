package main

import (
	"crypto/sha256"
	"time"
)

/*
	v1.1
	* 定义区块结构
	* 前区块哈希
	* 当前区块哈希
	* 数据

	v2.0
	*版本号	version
	*梅克尔根 merkelRoot
	*时间戳 Timestamp
	*难度值 difficulty
	*随机数 nonce
*/
type Block struct {
	Version    uint64 // 版本
	MerkelRoot []byte
	Timestamp  uint64
	Difficulty uint64
	Nonce      uint64
	PrevHash   []byte //前区块hash
	Hash       []byte //当前区块hash
	Data       []byte //数据
}

//    * 生成哈希
func (block *Block) SetHash() {
	// 获取区块信息
	blockInfo := []byte{}
	blockInfo = append(blockInfo, Uint64ToByte(block.Version)...)
	blockInfo = append(blockInfo, block.MerkelRoot...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Timestamp)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Difficulty)...)
	blockInfo = append(blockInfo, Uint64ToByte(block.Nonce)...)
	blockInfo = append(blockInfo, block.PrevHash...)
	blockInfo = append(blockInfo, block.Data...)
	// 生成hash
	hash := sha256.Sum256(blockInfo)

	block.Hash = hash[:]
}

// uint64转byte
func Uint64ToByte(num uint64) []byte {
	return []byte{}
}

//   * 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    00,
		MerkelRoot: []byte{},
		Timestamp:  uint64(time.Now().Unix()),
		Difficulty: 0,
		Nonce:      0,
		PrevHash:   prevBlockHash,
		Hash:       []byte{},
		Data:       []byte(data),
	}

	block.SetHash()

	return &block
}

//  创世区块
func GenesisBlock() *Block {
	return NewBlock("这是一个创世区块", []byte{})
}
