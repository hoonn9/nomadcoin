package blockchain

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/hoonn9/nomadcoin/db"
	"github.com/hoonn9/nomadcoin/utils"
)


type blockchain struct {
	NewestHash 	string	`json:"newestHash"`
	Height 		int		`json:"height"`
	CurrentDifficulty	int	`json:"currentDifficulty"`
	m									sync.Mutex
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

func (b *blockchain) AddBlock() *Block {
	block := createBlock(b.NewestHash, b.Height + 1, getDifficulty(b))
	b.NewestHash = block.Hash
	b.Height = block.Height
	b.CurrentDifficulty = block.Difficulty
	persistBlockchain(b)

	return block
}


func persistBlockchain(b *blockchain) {
	db.SaveCheckpoint(utils.ToBytes(b))
}


func Blocks(b *blockchain) []*Block {
	b.m.Lock()
	defer b.m.Unlock()

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

func Txs(b *blockchain) []*Tx {
	var txs []*Tx

	for _, block := range Blocks(b) {
		txs = append(txs, block.Transactions...)
	}

	return txs
}

func FindTx(b *blockchain, targetID string) *Tx {
	for _, tx := range Txs(b) {
		if tx.ID == targetID {
			return tx
		}
	}
	return nil
}

func recalculateDiffculty(b *blockchain) int {
	allBlocks := Blocks(b)
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

func getDifficulty(b *blockchain) int {
	if b.Height == 0 {
		return defaultDifficulty
	} else if b.Height % difficultyInterval == 0 {
		return recalculateDiffculty(b);
	} else {
		return b.CurrentDifficulty
	}
}


func UTxOutsByAddress(address string, b *blockchain) []*UTxOut {
	var uTxOuts []*UTxOut
	creatorTxs := make(map[string]bool)

	// 이미 사용한 output인지 확인
	for _, block := range Blocks(b) {
		for _, tx := range block.Transactions {
			for _, input := range tx.TxIns {
				if input.Signature == "Coinbase" {
					break
				}

				if FindTx(b, input.TxID).TxOuts[input.Index].Address == address {
					creatorTxs[input.TxID] = true
				}
			}
			for index, output := range tx.TxOuts {
				if output.Address == address {
					if _, ok := creatorTxs[tx.ID]; !ok {
						uTxOut := &UTxOut{tx.ID, index, output.Amount}
						// mempool에 올려져있는 output 이라면 사용 불가
						if !isOnMempool(uTxOut) {
							uTxOuts = append(uTxOuts, uTxOut)
						}
					}
				}
			}
		}
	}

	return uTxOuts
}

func BalanceByAddress(address string, b *blockchain) int {
	txOuts := UTxOutsByAddress(address, b)
	var amount int
	
	for _, txOut := range txOuts {
		amount += txOut.Amount
	}

	return amount
}

func Blockchain() *blockchain {
	once.Do(func() {
		b = &blockchain{
			Height: 0,
		}

		// 이전의 Block들 복구
		checkpoint := db.Checkpoint()

		// checkpoint not found
		if checkpoint == nil {
			b.AddBlock()
		} else {
			// search for checkpoint on the db
			// restore
			b.restore(checkpoint)
		}
	})

	return b
}

func Status(b *blockchain, rw http.ResponseWriter) {
	b.m.Lock()
	defer b.m.Unlock()

	utils.HandleErr(json.NewEncoder(rw).Encode(b))
}

func (b *blockchain) Replace(newBlocks []*Block) {
	b.m.Lock()
	defer b.m.Unlock()

	b.CurrentDifficulty = newBlocks[0].Difficulty
	b.Height = len(newBlocks)
	b.NewestHash = newBlocks[0].Hash

	persistBlockchain(b)
	db.EmptyBlocks()

	for _, block := range newBlocks {
		persistBlock(block)
	}
}

func (b *blockchain) AddPeerBlock(block *Block) {
	b.m.Lock()
	defer b.m.Unlock()

	b.Height += 1
	b.CurrentDifficulty = block.Difficulty
	b.NewestHash = block.Hash

	persistBlockchain(b)
	persistBlock(block)

	// mempool

}