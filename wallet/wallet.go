package wallet

import (
	"crypto/ecdsa"
	"os"
)

type wallet struct {
	privateKey	*ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	// exist check file
	_, err := os.Stat("nomadcoin.wallet")

	// file not exist error 인지 확인
	return os.IsExist(err)
}

func Wallet() *wallet {
	if w == nil {
		// has a wallet already
		// yes: restore from file
		// n o: create private key, save to file
		if hasWalletFile() {

		}
	}

	return w
}

func Start() {

}