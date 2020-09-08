package main

import (
	"fmt"

	"github.com/juanhuttemann/gochain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("First Block After Genesis")
	chain.AddBlock("Second Block After Genesis")
	chain.AddBlock("Third Block After Genesis")

	for _, block := range chain.Blocks {
		fmt.Println("Prev Hash:", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Println("Hash", block.Hash)
	}
}
