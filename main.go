package main

import (
	"github.com/hoonn9/nomadcoin/blockchain"
	"github.com/hoonn9/nomadcoin/cli"
	"github.com/hoonn9/nomadcoin/db"
)


func main() {
	defer db.Close()
	
	blockchain.Blockchain()
	cli.Start()
	// blockchain.Blockchain()
	// blockchain.Blockchain().AddBlock("First")
	// blockchain.Blockchain().AddBlock("Second")
	// blockchain.Blockchain().AddBlock("Third")

}
