package utils

import (
	"bytes"
	"encoding/gob"
	"log"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}


// interface{} any argument
// gob package  byte encoding, decoding
// Buffer  byte를 넣는 공간
func ToBytes( i interface{}) []byte {
	var aBuffer bytes.Buffer
	HandleErr(gob.NewEncoder(&aBuffer).Encode(i))
	return aBuffer.Bytes()
}