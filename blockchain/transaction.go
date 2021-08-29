package blockchain

import (
	"errors"
	"time"

	"github.com/hoonn9/nomadcoin/utils"
)

// 1. Coinbase 방식

const (
	minerReward int = 10
)

type mempool struct {
	Txs []*Tx
}

// 메모리에 존재. BlockChain 처럼 싱글톤 패턴, 초기화할 필요 없음
var Mempool *mempool = &mempool{}

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
	Owner 	string 	`json:"owner"`
	Amount	int		`json:"amount"`
}

type TxOut struct {
	Owner	string 	`json:"owner"`
	Amount	int		`json:"amount"`
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

func makeTx(from, to string, amount int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough money")
	}
	// 송금 양 amount 만큼 txIn 생성
	var txIns []*TxIn
	var txOuts []*TxOut
	total := 0
	oldTxOuts := Blockchain().TxOutsByAddress(from)

	// 같거나 커질 때까지
	for _, txOut := range oldTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{txOut.Owner, txOut.Amount}
		txIns = append(txIns, txIn)
		total += txOut.Amount
	}

	// 잔돈
	change := total - amount
	if change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}

	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)

	tx := &Tx {
		Id: "",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()

	return tx, nil
 }

// from 은 지갑에서 받아옴
func (m *mempool) AddTx(to string, amount int) error {
	tx, err := makeTx("nico", to, amount)
	if err != nil {
		return err
	}

	m.Txs = append(m.Txs, tx)
	return nil
}