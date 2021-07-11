package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data string
	hash string
	prevHash string
}

// 단방향 함수
// 입력 값으로 얻은 출력값을 출력값을 통해 입력값을 얻을 수 없음

/*
	하나의 hash만 변해도 전체가 변함. 무효화 된다.

	B1
		b1Hash = (data + "")
	B2
		b2Hash = (data + b1Hash)
	B3
		b3Hash = (data + b2Hash)
*/

func main() {
	// for _, aByte := range "Genesis Block" {
	// 	fmt.Printf("%b\n",aByte)
	// }

	genesisBlock := block{"Genesis Block", "", ""}
	fmt.Println(genesisBlock)
	// slice of byte

	/*
		SHA256
	*/

	hash := sha256.Sum256([]byte(genesisBlock.data + genesisBlock.hash))
	// 16진수로 format base16 => string으로 변환
	hexHash := fmt.Sprintf("%x",hash)
	genesisBlock.hash = hexHash
	fmt.Println(genesisBlock)
}
