package main

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/hoonn9/nomadcoin/db"
)


func main() {
	defer db.Close()
	
	// blockchain.Blockchain()
	// cli.Start()
	// blockchain.Blockchain()
	// blockchain.Blockchain().AddBlock("First")
	// blockchain.Blockchain().AddBlock("Second")
	// blockchain.Blockchain().AddBlock("Third")

	difficulty := 6
	target := strings.Repeat("0", difficulty)
	nonce := 1
	for {
		hash := fmt.Sprintf("%x", sha256.Sum256([]byte("hello" + fmt.Sprint(nonce))))
		fmt.Printf("Hash:%s\nTarget:%s\nNonce%d\n\n", hash, target, nonce)
		// 문자열이 a 값으로 시작하는지 확인 strings HasPrefix
		if strings.HasPrefix(hash, target) {
			return
		} else {
			nonce++
		}
	}
}
