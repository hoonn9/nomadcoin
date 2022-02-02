package blockchain

import (
	"crypto/ecdsa"
	"errors"
	"sync"
	"time"

	"github.com/hoonn9/nomadcoin/utils"
	"github.com/hoonn9/nomadcoin/wallet"
)

type myWallet interface {
	PrivateKey() *ecdsa.PrivateKey
	Address() string 
	Sign(payload string) string
}

var MyWallet myWallet = wallet.Wallet{}

// 1. Coinbase 방식

const (
	minerReward int = 10
)

type mempool struct {
	Txs map[string]*Tx
	m		sync.Mutex
}

// 메모리에 존재. BlockChain 처럼 싱글톤 패턴, 초기화할 필요 없음
var m *mempool
var memOnce sync.Once

func Mempool() *mempool {
	memOnce.Do(func() {
		m = &mempool{
			Txs: make(map[string]*Tx),
		}
	})
	return m
}

type Tx struct {
	ID			string		`json:"id"`	
	Timestamp	int   		`json:"timestamp"`	// 거래 발생 시간
	TxIns		[]*TxIn		`json:"txIns"`		// 입력값
	TxOuts		[]*TxOut	`json:"txOuts"`		// 출력값
}

type TxIn struct {
	TxID 	string	`json:"txId"`
	Index 	int		`json:"index"`
	Signature 	string 	`json:"signature"`
}

type TxOut struct {
	Address	string 	`json:"address"`
	Amount	int		`json:"amount"`
}

// Unspend Transaction
type UTxOut struct {
	TxID	string
	Index	int
	Amount	int
}

var ErrorNoMoney = errors.New("not enough amount")
var ErrorNotValid = errors.New("transaction invaild")

func (tx *Tx) getId() {
	tx.ID = utils.Hash(tx)
}

func (tx *Tx) sign() {
	for _, txIn := range tx.TxIns {
		txIn.Signature = MyWallet.Sign(tx.ID)
	}
}

func validate(tx *Tx) bool {
	valid := true

	// input에 참조되는 output의 소유를 증명

	for _, txIn := range tx.TxIns {
		// txIn.TXID => input으로 쓰이는 output을 만든 transaction ID
		prevTx := FindTx(Blockchain(), txIn.TxID)

		if prevTx == nil {
			valid = false
			break
		}

		address := prevTx.TxOuts[txIn.Index].Address
		valid = wallet.Verify(txIn.Signature, tx.ID, address)

		if !valid {
			break
		}
	}

	return valid
}

func isOnMempool(uTxOut *UTxOut) bool {
	exists := false

	// label => 이중 for loop 에서 break 하는 방법
	Outer:
		for _, tx := range Mempool().Txs {
			for _, input := range tx.TxIns {
				if input.TxID == uTxOut.TxID && input.Index == uTxOut.Index {
					exists = true
					break Outer
				}
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

func makeTx(to string, amount int) (*Tx, error) {
	if BalanceByAddress(MyWallet.Address(), Blockchain()) < amount {
		return nil, ErrorNoMoney
	}

	var txOuts []*TxOut
	var txIns []*TxIn

	total := 0
	uTxOuts := UTxOutsByAddress(MyWallet.Address(), Blockchain())

	for _, uTxOut := range uTxOuts {
		if total >= amount {
			break
		}
		txIn := &TxIn{uTxOut.TxID, uTxOut.Index, MyWallet.Address()}
		txIns = append(txIns, txIn)
		total += uTxOut.Amount
	}

	if change := total -  amount; change != 0 {
		changeTxOut := &TxOut{MyWallet.Address(), change}
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
	tx.sign()
	valid := validate(tx)

	if !valid {
		return nil, ErrorNotValid
	}

	return tx, nil
 }

// from 은 지갑에서 받아옴
func (m *mempool) AddTx(to string, amount int) (*Tx,error) {
	tx, err := makeTx(to, amount)
	if err != nil {
		return nil, err
	}

	m.Txs[tx.ID] =  tx
	return tx, nil
}

func (m *mempool) txToConfirm() []*Tx {
	coinbase := makeCoinbaseTx(MyWallet.Address())

	var txs []*Tx

	for _, tx := range m.Txs {
		txs = append(txs, tx)
	}

	txs = append(txs, coinbase)
	m.Txs = make(map[string]*Tx)

	return txs
}

func (m *mempool) AddPeerTx(tx *Tx) {
	m.m.Lock()
	defer m.m.Unlock()

	m.Txs[tx.ID] = tx
}