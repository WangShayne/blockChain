package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 定义工作量证明结构
type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// 创建工作量证明
func NewProofOfWord(block *Block) *ProofOfWork {
	pow := ProofOfWork{
		block: block,
	}

	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	targetInt := big.Int{}
	targetInt.SetString(targetStr, 16)
	pow.target = &targetInt
	return &pow
}

// 提供不断计算的hash函数 Run()
func (pow *ProofOfWork) Run() ([]byte, uint64) {
	//TODO

	var nonce uint64
	block := pow.block
	hash := [32]byte{}

	for {
		tmp := [][]byte{
			Uint64ToByte(block.Version),
			block.MerkelRoot,
			Uint64ToByte(block.Timestamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.PrevHash,
			block.Data,
		}

		bookInfo := bytes.Join(tmp, []byte{})
		// 做哈希运算
		hash = sha256.Sum256(bookInfo)
		// 与pow的target做比较
		tmpInt := big.Int{}
		tmpInt.SetBytes(hash[:])

		// 比较hash和目标hash
		// -1 x < y
		// 0 x == y
		// 1 x > y

		if tmpInt.Cmp(pow.target) == -1 {
			fmt.Printf("挖矿成功,hash值为:%x,随机数为:%d \n", hash, nonce)
			break
		} else {
			nonce++
		}

	}

	return hash[:], nonce
}

//  提供一个校验函数 IsValid()
