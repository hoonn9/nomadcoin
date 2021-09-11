package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"os"

	"github.com/hoonn9/nomadcoin/utils"
)

const (
	fileName string = "nomadcoin.wallet"
)

type wallet struct {
	privateKey	*ecdsa.PrivateKey
}

var w *wallet

func hasWalletFile() bool {
	// exist check file
	_, err := os.Stat(fileName)

	// file not exist error 인지 확인
	return os.IsExist(err)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	utils.HandleErr(err)

	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)

	utils.HandleErr(err)

	// variable, err 반환할 때 new variable이 오면 err := 로 가능하지만 아니면 err = 로 갱신 해줘야 함
	err = os.WriteFile(fileName, bytes, 0644)

	utils.HandleErr(err)
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		// has a wallet already

		// yes: restore from file
		if hasWalletFile() {
			
		} else { // no: create private key, save to file
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		
	}

	return w
}

func Start() {

}