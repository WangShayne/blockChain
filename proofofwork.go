package main

import "math/big"

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
	return []byte("helloww"), 10
}

//  提供一个校验函数 IsValid()
