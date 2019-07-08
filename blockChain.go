package main


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
