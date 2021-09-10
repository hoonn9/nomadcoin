package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/hex"
	"fmt"

	"github.com/hoonn9/nomadcoin/utils"
)

/*
	sign, verify


	1) hash the msg
	"message" -> hash(x) -> "hashed_message"

	2) generate key pair
	private key, public key (using go, save private key to a file)

	3) sign the hash
	("hashed_message" + private key) -> "signature"

	4) verify
	("hashed_message" + "signature" + public key) -> true | false
*/

const (
	// signature = "86f77bb4b933f543179ba71fdf9a4fea2f628d21dcfbc38a8de23da351c55db3260ec43a0d02f1812e99dca6f04f0d62b2d29cf93ca9a5089bc5d81a212b9765%"
	hashedMessage = "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	// privateKey = "30770201010420436c85b83ee3022a9cf7e0c1fa42f4dd745ffb2ffbb309e2cf39b3950fa16b67a00a06082a8648ce3d030107a144034200049c13419b37e1e0039433c8f25e4b7118e02d35ee076c952eb53103d072b6f68dfc112df9f4e7fd4b05cf7ef2e30850f9c23a090e31c293c15a0568700eab3d47"
)

func Start() {
	// public key 타원 곡선 상의 x, y 값

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	// privateKey를 bytes 로 변환
	keyAsBytes, err := x509.MarshalECPrivateKey(privateKey)

	fmt.Printf("%x\n\n",keyAsBytes)

	utils.HandleErr(err)

	fmt.Println(hashedMessage)

	hashAsBytes, err := hex.DecodeString(hashedMessage)

	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	// r, s bytes len가 32로 동일 => 합쳐서 32로 나누면 다시 restore 가능
	fmt.Println(r.Bytes(), s.Bytes())

	// slice[], element
	// anotherSlice... => element 로 뺴오는 방법
	signature := append(r.Bytes(), s.Bytes()...)

	fmt.Printf("%x", signature)

	utils.HandleErr(err)

	
}