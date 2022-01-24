package utils

import (
	"encoding/hex"
	"errors"
	"fmt"
	"reflect"
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

func TestToBytes(t *testing.T) {
	s := "test"
	b := ToBytes(s)
	kind := reflect.TypeOf(b).Kind() 
	if kind != reflect.Slice {
		t.Errorf("ToBytes should return a slice of bytes got %s", kind)
	}
}

func TestSplitter(t *testing.T) {
	type test struct {
		input string
		sep 	string
		index	int
		output string
	}

	tests := []test{
		{input: "0:6:0", sep: ":", index: 1, output: "6"},
		{input: "0:6:0", sep: ":", index: 10, output: ""},
		{input: "0:6:0", sep: "/", index: 0, output: "0:6:0"},
	}

	for _, tc := range tests {
		got := Splitter(tc.input, tc.sep, tc.index)
		if got != tc.output  {
			t.Errorf("Expected %s and got %s", tc.output, got)
		}
	}
}

func TestHandleErr(t *testing.T) {
	oldLogFn := logFn
	defer func() {
		logFn = oldLogFn
	}()
	called := false
	logFn = func(v ...interface{}) {
		called = true
	}
	err := errors.New("test")
	HandleErr(err) 

	if !called {
		t.Error("HandleErr should call fn")
	}
}