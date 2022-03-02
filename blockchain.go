package main

type Blockchain struct {
	blocks []*Block
}

// 创建一个带有创始区块的区块链。
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func NewGenesisBlock() *Block {
	return NewBlock("0.0-beta0", "Genesis Block", []byte{})
}

// TODO: Version和Nonce的值现在是随便继承了之前的。
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(prevBlock.Version, data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
