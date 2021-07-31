package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/hoonn9/nomadcoin/db"
	"github.com/hoonn9/nomadcoin/utils"
)

// n<difficulty>개의 0으로 시작하는 해쉬 값을 찾는다.
const difficulty int = 2

// Nonce  채굴자들이 유일하게 바꿀 수 있는 값 (1회성 값)

type Block struct {
	Data 		string 	`json:"data"`
	Hash 		string	`json:"hash"`
	PrevHash 	string	`json:"prevHash,omitempty"`
	Height 		int		`json:"height"`
	Difficulty	int		`json:"difficulty"`
	Nonce		int		`json:"nonce"`
}

var ErrNotFound = errors.New("block not found")


func (b *Block) persist() {
	db.SaveBlock(b.Hash, utils.ToBytes(b))
}

func (b *Block) restore(data []byte) {
	utils.FromBytes(b, data)
}

func FindBlock(hash string) (*Block, error){
	blockBytes := db.Block(hash)
	if blockBytes == nil {
		return nil, ErrNotFound
	}
	block := &Block{}
	block.restore(blockBytes)
	
	return block, nil
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

