package blockchain

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"

	"github.com/hoonn9/nomadcoin/db"
	"github.com/hoonn9/nomadcoin/utils"
)


type blockchain struct {
	NewestHash 	string	`json:"newestHash"`
	Height 		int		`json:"height"`
}

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	decoder := gob.NewDecoder(bytes.NewReader(data))
	// decode only pointer
	utils.HandleErr(decoder.Decode(b))
}

func (b *blockchain) persist() {
	db.SaveBlockchain(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.persist()
}


func Blockchain() *blockchain {
	// only once execute
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}

			fmt.Printf("Newest Hash: %s\nHeight: %d\n", b.NewestHash, b.Height)
			
			checkpoint := db.Checkpoint()

			// checkpoint not found
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				// search for checkpoint on the db
				// restore
				fmt.Println("Restoring...")

				b.restore(checkpoint)
			}
			
		
		})
	}
	fmt.Printf("Newest Hash: %s\nHeight: %d\n", b.NewestHash, b.Height)

	return b
}

