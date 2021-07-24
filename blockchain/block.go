package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"

	"github.com/hoonn9/nomadcoin/db"
	"github.com/hoonn9/nomadcoin/utils"
)

type Block struct {
	Data 	string 	`json:"data"`
	Hash 	string	`json:"hash"`
	PrevHash string	`json:"prevHash,omitempty"`
	Height 	int		`json:"height"`
}

// gob package  byte encoding, decoding
// Buffer  byte를 넣는 공간
func (b *Block) toBytes() []byte {
	var blockBuffer bytes.Buffer
	utils.HandleErr(gob.NewEncoder(&blockBuffer).Encode(b))
	return blockBuffer.Bytes()
}

func (b *Block) persist() {
	db.SaveBlock(b.Hash, b.toBytes())
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data: data,
		Hash: "",
		PrevHash: prevHash,
		Height: height,
	}
	payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))
	block.persist()
	return block
}