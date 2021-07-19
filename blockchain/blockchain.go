package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data string
	Hash string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}


// 단방향 함수
// 입력 값으로 얻은 출력값을 출력값을 통해 입력값을 얻을 수 없음

/*
	하나의 hash만 변해도 전체가 변함. 무효화 된다.

	B1
		b1Hash = (data + "")
	B2
		b2Hash = (data + b1Hash)
	B3
		b3Hash = (data + b2Hash)
*/


/*
	singleton: application 내 어디서든 단 하나의 instance를 공유하는 패턴
*/

var b *blockchain
var once sync.Once

/*
	sync package
	동기적 실행 관리

	코드를 thread가 몇 개이던, goroutine 이던 단 한번만 실행 => sync.Once
*/

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks - 1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	// only once execute
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}