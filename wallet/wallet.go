package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
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

func Start() {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	// public key 타원 곡선 상의 x, y 값

	utils.HandleErr(err)

	fmt.Println("Private Key", privateKey.D)
	fmt.Println("Public Key. x, y", privateKey.X, privateKey.Y)

	message := "i love you"

	hasedMessage :=  utils.Hash(message)

	hashAsBytes, err := hex.DecodeString(hasedMessage)

	utils.HandleErr(err)

	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashAsBytes)

	utils.HandleErr(err)

	fmt.Printf("R:%d\nS:%d", r, s)
}