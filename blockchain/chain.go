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
	CurrentDifficulty	int	`json:"currentDifficulty"`
}

const (
	defaultDifficulty 	int = 2
	difficultyInterval 	int = 5
	blockInterval		int = 2
	allowedRange		int = 2
)

var b *blockchain
var once sync.Once

func (b *blockchain) restore(data []byte) {
	utils.FromBytes(b, data)
}

func (b *blockchain) persist() {
	db.SaveCheckpoint(utils.ToBytes(b))
}

func (b *blockchain) AddBlock(data string) {
	block := createBlock(data, b.NewestHash, b.Height + 1)
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
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

func (b *blockchain) recalculateDiffculty() int {
	allBlocks := b.Blocks()
	newestBlock := allBlocks[0]
	recalculatedBlock := allBlocks[difficultyInterval - 1]

	// now.unix는 초단위 따라서 60으로 나눈다.
	actualTime := (newestBlock.Timestamp / 60) - (recalculatedBlock.Timestamp / 60)

	// 소요시간 기대값. 5개당 검사. mine하는데 2분 => 총 10분
	// 기대값보다 작으면 난이도를 높이고 크면 난이도를 낮춘다.
	expectedTime := difficultyInterval * blockInterval
	
	// 소요시간 기대값에 근접 (allowedRange 만큼) 이면 난이도 유지
	if actualTime < (expectedTime - allowedRange) {
		return b.CurrentDifficulty + 1
	} else if actualTime > (expectedTime + allowedRange) {
		return b.CurrentDifficulty - 1
	}
	return b.CurrentDifficulty
}

func (b *blockchain) difficulty() int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		return b.recalculateDiffculty();
	} else {
		return b.CurrentDifficulty
	}
}

func Blockchain() *blockchain {
	// only once execute
	if b == nil {
		once.Do(func() {
			b = &blockchain{
				Height: 0,
			}

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

