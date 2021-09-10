package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/hoonn9/nomadcoin/utils"
)



const (
	signature 		string	= "86f77bb4b933f543179ba71fdf9a4fea2f628d21dcfbc38a8de23da351c55db3260ec43a0d02f1812e99dca6f04f0d62b2d29cf93ca9a5089bc5d81a212b9765%"
	hashedMessage 	string	= "1c5863cd55b5a4413fd59f054af57ba3c75c0698b3851d70f99b8de2d5c7338f"
	privateKey 		string	= "30770201010420436c85b83ee3022a9cf7e0c1fa42f4dd745ffb2ffbb309e2cf39b3950fa16b67a00a06082a8648ce3d030107a144034200049c13419b37e1e0039433c8f25e4b7118e02d35ee076c952eb53103d072b6f68dfc112df9f4e7fd4b05cf7ef2e30850f9c23a090e31c293c15a0568700eab3d47"
)

func Start() {
	
	// restore

	// private key restore
	// 16인수 string인지 check
	privateBytes, err := hex.DecodeString(privateKey)
	utils.HandleErr(err)

	restoreKey, err := x509.ParseECPrivateKey([]byte(privateBytes))
	utils.HandleErr(err)

	sigBytes, err := hex.DecodeString(signature)

	rBytes := sigBytes[:len(sigBytes)/2]
	sBytes := sigBytes[len(sigBytes)/2:]

	fmt.Println(restoreKey)

	var bigR, bigS = big.Int{}, big.Int{}

	bigR.SetBytes(rBytes)
	bigS.SetBytes(sBytes)

	fmt.Println(bigR, bigS)
}