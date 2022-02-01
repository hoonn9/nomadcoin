package main

import (
	"github.com/hoonn9/nomadcoin/cli"
	"github.com/hoonn9/nomadcoin/db"
)


func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
