// Package utils contains function to be used across the application.
package utils

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}




// interface{} any argument
// gob package  byte encoding, decoding
// Buffer  byte를 넣는 공간
func ToBytes(i interface{}) []byte {
	var aBuffer bytes.Buffer
	HandleErr(gob.NewEncoder(&aBuffer).Encode(i))
	return aBuffer.Bytes()
}

// FromBytes takes an interface and data and then will encode the data to the interface.
func FromBytes(i interface{}, data []byte)  {
	encoder := gob.NewDecoder(bytes.NewReader(data))
	HandleErr(encoder.Decode(i))
}

// Hash takes an interface, hashes it and returns the hex encoding of the hash.
func Hash(i interface{}) string {
	// %v 기본 format
	s := fmt.Sprintf("%v", i)
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func Splitter(s,sep string, i int) string {
	r := strings.Split(s, sep)

	if len(r) - 1 < i {
		return ""
	}
	return r[i]
}


func ToJSON(i interface{}) []byte {
	r, err := json.Marshal(i)
	HandleErr(err)

	return r
}