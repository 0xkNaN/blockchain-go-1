package main

import (
	"fmt"

	"home/hassen/workspaces_blockchain/chain-go-1/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()

	chain.AddBlock("Block Test - 1")
	chain.AddBlock("Block Test - 2")
	chain.AddBlock("Block Test - 3")

	for _, block := range chain.blocks {
		fmt.Printf("Data  ::  %s\n", block.Data)
		fmt.Printf("cHash ::  %x\n", block.Hash)
		fmt.Printf("pHash ::  %x\n\n\n", block.PrevHash)
	}
}
