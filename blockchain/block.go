package blockchain

import (
	"errors"
	"strings"
	"time"

	"github.com/hoonn9/nomadcoin/db"
	"github.com/hoonn9/nomadcoin/utils"
)

// n<difficulty>개의 0으로 시작하는 해쉬 값을 찾는다.
// Nonce  채굴자들이 유일하게 바꿀 수 있는 값 (1회성 값)

type Block struct {
	Data 		string 	`json:"data"`
	Hash 		string	`json:"hash"`
	PrevHash 	string	`json:"prevHash,omitempty"`
	Height 		int		`json:"height"`
	Difficulty	int		`json:"difficulty"`
	Nonce		int		`json:"nonce"`
	Timestamp	int		`json:"timestamp"`
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

func (b *Block) mine() {
	target := strings.Repeat("0", b.Difficulty)
	for {
		b.Timestamp = int(time.Now().Unix())
		hash := utils.Hash(b)
		if strings.HasPrefix(hash, target) {
			b.Hash = hash
			break
		} else {
			b.Nonce++
		}
	}
}

func createBlock(data string, prevHash string, height int) *Block {
	block := &Block{
		Data: data,
		Hash: "",
		PrevHash: prevHash,
		Height: height,
		Difficulty: Blockchain().difficulty(),
		Nonce: 0,
	}
	// hash 생성에 값들을 나열하면서 붙이는 방법이 좋지 못함.
	// block 자체를 Hash
	// payload := block.Data + block.PrevHash + fmt.Sprint(block.Height)
	// block.Hash = fmt.Sprintf("%x", sha256.Sum256([]byte(payload)))

	block.mine()
	block.persist()
	return block
}

