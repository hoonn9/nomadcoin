package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"io/fs"
	"math/big"
	"os"

	"github.com/hoonn9/nomadcoin/utils"
)

const (
	fileName string = "nomadcoin.wallet"
)

type fileLayer interface {
	hasWalletFile() bool
	writeFile(name string, data []byte, perm fs.FileMode) error
	readFile(name string) ([]byte, error)
}

type layer struct{}

func (layer) hasWalletFile() bool {
	_, err := os.Stat(fileName)
	return !os.IsNotExist(err)
}

func (layer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (layer) readFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

var files fileLayer = layer{}

type Wallet struct {}

var w *Wallet

func (Wallet) PrivateKey() *ecdsa.PrivateKey {
	if files.hasWalletFile() {
		return restoreKey()
	} 
	key := createPrivateKey()
	persistKey(key)
	return key
}

func (w Wallet) Address() string {
	return aFromK(Wallet.PrivateKey(w))
}

func (w Wallet) Sign(payload string) string {
	return sign(payload, &w)
}

func createPrivateKey() *ecdsa.PrivateKey {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	utils.HandleErr(err)

	return privateKey
}

func persistKey(key *ecdsa.PrivateKey) {
	bytes, err := x509.MarshalECPrivateKey(key)
	utils.HandleErr(err)
	err = files.writeFile(fileName, bytes, 0644)
	// variable, err 반환할 때 new variable이 오면 err := 로 가능하지만 아니면 err = 로 갱신 해줘야 함
	utils.HandleErr(err)
}


// named return => short function에서 사용
func restoreKey() (key *ecdsa.PrivateKey) {
	keyAsBytes, err := files.readFile(fileName)
	utils.HandleErr(err)

	key, err = x509.ParseECPrivateKey(keyAsBytes)
	utils.HandleErr(err)

	return
}

func restoreBigInts(payload string) (*big.Int, *big.Int, error) {
	bytes, err := hex.DecodeString(payload)

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

func sign(payload string, w *Wallet) string {
	payloadAsBytes, err := hex.DecodeString(payload)
	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, w.PrivateKey(), payloadAsBytes)
	utils.HandleErr(err)

	return encodeBigInts(r.Bytes(), s.Bytes())
}

func Verify(signature, payload, address string) bool {
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

func InitWallet() {
	if w == nil {
		w = &Wallet{}
	}
}