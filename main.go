package main

import (
	"fmt"

	"github.com/hoonn9/nomadcoin/blockchain"
)




func main() {
	chain := blockchain.GetBlockchain()
	fmt.Println(chain)
}
