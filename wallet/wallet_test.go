package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"io/fs"
	"reflect"
	"testing"
)

const (
	testKey = "3077020101042075b8f2cde7b87b7ee2d7e99a3a2c6d2791dc0a1004bd842c14308b897f51bad0a00a06082a8648ce3d030107a144034200049bc45b24483f6cb358e10faad945ef67b7394ba214b882f6eb32d78f0b29181c7e868a260c9ca0c40f1e7fcf676fb517212562fb1ce4cefb9e0184970b0c0eab"
	testPayload = "0000bf1c1471cae6688756bad32c6b958975cc25ccd623fd82cfe4cc0d05f359"
	testSignature = "3750e4b72fcb7057be14b754f781d09ffc0331e53d6af6ed4429b1b89958f5fe9217852f2d87640fd0f32e0fe134dee0ccb221626620b1bc6375151c3c298f55"
)

type fakeLayer struct{
	fakeHasWalletFile func() bool
}


func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()
}

func (fakeLayer) writeFile(name string, data []byte, perm fs.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	// read file의 목적은 bytes를 리턴
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool {return false},
		}

		w :=  Wallet()
		if reflect.TypeOf(w) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})

	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool {return true},
		}
		// created test에서 생성된 wallet 비워주기
		w = nil		
		w :=  Wallet()
		if reflect.TypeOf(w) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
}


func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)

	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok bool
	}

	tests := []test {
		{testPayload, true},
		{"0000bf1c1471cae6688756bad32c6b958975cc25ccd623fd82cfe4cc0d05f352", false},
	}

	for _, tc := range tests {

		w := makeTestWallet()
		ok := Verify(testSignature, testPayload, w.Address)

		if ok != tc.ok {
			t.Error("Verify could not verify testSignature and Payload")
		}
	}
}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")

	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex")
	}
}

func TestSign(t *testing.T) {
	signature := Sign(testPayload, makeTestWallet())

	_, err := hex.DecodeString(signature)

	if err != nil {
		t.Errorf("Sign() should return a hex encoded string, got %s", signature)
	}
}