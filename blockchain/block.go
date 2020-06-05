package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type HashType string

type Block struct {
	header    BlockHeader
	txs       string
	txCounter int
	hashCurr  HashType // 当前区块 hash
}

type BlockHeader struct {
	version        int
	hashPrevBlock  HashType
	hashMerkleRoot HashType
	time           time.Time
	bits           int
	nonce          int
}

func NewBlock(hashPrevBlock HashType, txs string) *Block {
	b := &Block{
		header: BlockHeader{
			hashPrevBlock: hashPrevBlock,
			time:          time.Now(),
		},
		txs:       txs,
		txCounter: 0,
		hashCurr: "",
	}
	b.SetHash()
	return b
}

func (h *BlockHeader) String() string {
	return fmt.Sprintf("%d%s%s%d%d%d",
		h.version,
		h.hashPrevBlock,
		h.hashMerkleRoot,
		h.time.UnixNano(),
		h.bits,
		h.nonce,
	)
}

func (b *Block) SetHash() *Block {
	b.hashCurr = HashType(fmt.Sprintf("%x", sha256.Sum256([]byte(b.header.String()))))
	return b
}
