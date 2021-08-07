package blockchain

import (
	"time"

	"github.com/hoonn9/nomadcoin/utils"
)

// 1. Coinbase 방식

const (
	minerReward int = 10
)

type Tx struct {
	Id			string		`json:"id"`	
	Timestamp	int   		`json:"timestamp"`	// 거래 발생 시간
	TxIns		[]*TxIn		`json:"txIns"`		// 입력값
	TxOuts		[]*TxOut	`json:"txOuts"`		// 출력값
}

func (t *Tx) getId() {
	t.Id = utils.Hash(t)
}

type TxIn struct {
	Owner 	string
	Amount	int
}

type TxOut struct {
	Owner	string
	Amount	int
}

func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{
			"Coinbase", minerReward,
		},
	}

	txOuts := []*TxOut{
		{
			address, minerReward,
		},
	}

	tx := Tx{
		Id: "",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()
	return &tx
}