package blockchain

import (
	"testing"

	"github.com/hoonn9/nomadcoin/utils"
)

func TestMakeTx(t *testing.T) {
	// w = &myWallet{}
	blocks := []*Block{
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: "x"},
		{PrevHash: ""},
	}
	fakeBlock := 0
	dbStorage = fakeDB{
		fakeFindBlock: func() []byte {
			defer func ()  {
				fakeBlock++
			}()
			return utils.ToBytes(blocks[fakeBlock])
		},
	}
	tx, err := makeTx("test", 10)

	if err != nil {
		t.Errorf("MakeTx should not error, got %s", err)
	}

	for _, txOut := range tx.TxOuts {
		if txOut.Address != "test" {
			t.Error("txOut address not match")
		}
	}
}