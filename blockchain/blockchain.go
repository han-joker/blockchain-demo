package blockchain

type Blockchain struct {
	hashLastBlock HashType
	blocks        map[HashType]*Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{
		hashLastBlock: "",
		blocks:        map[HashType]*Block{},
	}
	bc.AddBlock("Genesis Block.")

	return bc
}

func (bc *Blockchain) AddBlock(txs string) *Blockchain {
	b := NewBlock(bc.hashLastBlock, txs)
	bc.hashLastBlock = b.hashCurr
	bc.blocks[b.hashCurr] = b

	return bc
}
