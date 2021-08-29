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
	ID			string		`json:"id"`	
	Timestamp	int   		`json:"timestamp"`	// 거래 발생 시간
	TxIns		[]*TxIn		`json:"txIns"`		// 입력값
	TxOuts		[]*TxOut	`json:"txOuts"`		// 출력값
}

func (t *Tx) getId() {
	t.ID = utils.Hash(t)
}

type TxIn struct {
	TxID 	string	`json:"txId"`
	Index 	int		`json:"index"`
	Owner 	string 	`json:"owner"`
}

type TxOut struct {
	Owner	string 	`json:"owner"`
	Amount	int		`json:"amount"`
}

// Unspend Transaction
type UTxOut struct {
	TxID	string
	Index	int
	Amount	int
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false

	for _, tx := range Mempool.Txs {
		for _, input := range tx.TxIns {
			exists = input.TxID == uTxOut.TxID && input.Index == uTxOut.Index 
		}
	}

	return exists
}


func makeCoinbaseTx(address string) *Tx {
	txIns := []*TxIn{
		{
			"", -1,	"Coinbase",
		},
	}

	txOuts := []*TxOut{
		{
			address, minerReward,
		},
	}

	tx := Tx{
		ID: "",
		Timestamp: int(time.Now().Unix()),
		TxIns: txIns,
		TxOuts: txOuts,
	}
	tx.getId()
	return &tx
}

func makeTx(from, to string, amount int) (*Tx, error) {
	if Blockchain().BalanceByAddress(from) < amount {
		return nil, errors.New("not enough amount")
	}

	var txOuts []*TxOut
	var txIns []*TxIn

	total := 0
	uTxOuts := Blockchain().UTxOutsByAddress(from)

	for _, uTxOut := range uTxOuts {
		if total > amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, from}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}

	if change := total -  amount; change != 0 {
		changeTxOut := &TxOut{from, change}
		txOuts = append(txOuts, changeTxOut)
	}

	txOut := &TxOut{to, amount}
	txOuts = append(txOuts, txOut)

	tx := &Tx{
		ID: "",
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

func (m *mempool) txToConfirm() []*Tx {
	coinbase := makeCoinbaseTx("nico")

	// tx 개수의 제한을 두지 않음
	txs := m.Txs
	txs = append(txs, coinbase)
	m.Txs = nil

	return txs
}