package main

import (
	"github.com/hoonn9/nomadcoin/blockchain"
	"github.com/hoonn9/nomadcoin/cli"
)


func main() {
	blockchain.Blockchain()
	cli.Start()
	// blockchain.Blockchain()
	// blockchain.Blockchain().AddBlock("First")
	// blockchain.Blockchain().AddBlock("Second")
	// blockchain.Blockchain().AddBlock("Third")

}
