package main

import (
	"github.com/hoonn9/nomadcoin/cli"
	"github.com/hoonn9/nomadcoin/db"
	"github.com/hoonn9/nomadcoin/wallet"
)


func main() {
	defer db.Close()
	db.InitDB()
	wallet.InitWallet()
	cli.Start()
}
