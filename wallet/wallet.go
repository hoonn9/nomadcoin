package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/hoonn9/nomadcoin/utils"
)

const (
	fileName string = "nomadcoin.wallet"
)

type wallet struct {
	privateKey	*ecdsa.PrivateKey
	Address		string
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


// named return => short function에서 사용
func restoreKey() (key *ecdsa.PrivateKey) {
	keyAsBytes, err := os.ReadFile(fileName)
	utils.HandleErr(err)

	key, err = x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)

	return
}

func restoreBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	if err != nil {
		return nil, nil, err
	}

	firstHalfBytes := bytes[:len(bytes) / 2]
	secondHalfBytes := bytes[len(bytes) / 2:]

	bigA, bigB := big.Int{}, big.Int{}
	bigA.SetBytes(firstHalfBytes)
	bigB.SetBytes(secondHalfBytes)

	return &bigA, &bigB, nil
}

func encodeBigInts(a, b []byte) string {
	z := append(a, b...)
	return fmt.Sprintf("%x", z)
}

func aFromK(key *ecdsa.PrivateKey) string {
	return encodeBigInts(key.X.Bytes(), key.Y.Bytes())
}

func sign(payload string, w *wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, w.privateKey, payloadAsBytes)
	utils.HandleErr(err)

	return encodeBigInts(r.Bytes(), s.Bytes())
}

func verify(signature, payload, address string) bool {
	r, s, err := restoreBigInts(signature)
	utils.HandleErr(err)

	x, y, err := restoreBigInts(address)
	utils.HandleErr(err)

	// make publicKey
	publicKey := ecdsa.PublicKey{
		Curve: elliptic.P256(),
		X: x,
		Y: y,
	}
	
	payloadAsBytes, err := hex.DecodeString(payload)

	utils.HandleErr(err)

	ok := ecdsa.Verify(&publicKey, payloadAsBytes, r, s)
	return ok
}

func Wallet() *wallet {
	if w == nil {
		w = &wallet{}
		// has a wallet already

		// yes: restore from file
		if hasWalletFile() {
			w.privateKey = restoreKey()
		} else { // no: create private key, save to file
			key := createPrivateKey()
			persistKey(key)
			w.privateKey = key
		}
		w.Address = aFromK(w.privateKey)
	}

	return w
}