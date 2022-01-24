package utils

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hash := "1cb6fa80c3cf0c4449500e4358cd5cdab5cbc43e035563c65ba4513d0c3fd116"
	s := struct{Test string }{Test: "string"}


	t.Run("Hash is always same", func(t *testing.T) {
		x := Hash(s)

		if x != hash {
			t.Errorf("Expected %s, got %s", hash, x)
		}
	})

	t.Run("Hash is always encoded", func(t *testing.T) {
		x := Hash(s)

		_, err := hex.DecodeString(x)

		if err != nil {
			t.Error("Hash should be hex encoded")
		}
	})
}

func ExampleHash() {
	s := struct{Test string }{Test: "string"}
	x := Hash(s)
	fmt.Println(x)
	// Output: 1cb6fa80c3cf0c4449500e4358cd5cdab5cbc43e035563c65ba4513d0c3fd116

}