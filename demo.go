package main

import (
	"fmt"
	"github.com/han-joker/blockchain-demo/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()
	bc.AddBlock("TX 1")
	bc.AddBlock("TX 2")
	bc.AddBlock("TX 3")
	bc.AddBlock("TX 4")
	fmt.Println(bc)
}
