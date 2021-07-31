package blockchain

import (
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
	// decoder := gob.NewDecoder(bytes.NewReader(data))
	// decode only pointer
	// utils.HandleErr(decoder.Decode(b))
	utils.FromBytes(b, data)
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


func (b *blockchain) Blocks() []*Block {
	var blocks []*Block
	hashCursor := b.NewestHash

	// while(true) 종료는 break
	for {
		block, _ := FindBlock(hashCursor)
		blocks = append(blocks, block)

		if block.PrevHash != "" {
			hashCursor = block.PrevHash
		} else {
			// PrevHash 가 없으면 Genesis block
			break
		}
	}

	return blocks
}


func Blockchain() *blockchain {
	// only once execute
	if b == nil {
		once.Do(func() {
			b = &blockchain{"", 0}

			// 이전의 Block들 복구
			checkpoint := db.Checkpoint()

			// checkpoint not found
			if checkpoint == nil {
				b.AddBlock("Genesis Block")
			} else {
				// search for checkpoint on the db
				// restore
				b.restore(checkpoint)
			}
		})
	}

	fmt.Println(b.NewestHash)
	return b
}

