package main

import "github.com/hoonn9/nomadcoin/blockchain"


func main() {
	// cli.Start()
	blockchain.Blockchain().AddBlock("First")
	blockchain.Blockchain().AddBlock("Second")
	blockchain.Blockchain().AddBlock("Third")

}
