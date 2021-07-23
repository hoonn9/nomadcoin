package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hoonn9/nomadcoin/utils"
)

const port string = ":4000"


type URLDescription struct {
	URL 		string
	Method 		string
	Description string
}
/**
	Marshal 
	메모리형식으로 저장된 객체를 저장, 송신 가능하게 만들어 주는 것
	Unmarshal
	반대
*/
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
	}
	// slice to json(byte[])
	b, err := json.Marshal(data)
	utils.HandleErr(err)
	fmt.Println(b)
	fmt.Printf("%s", b)
}

func main() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on http://localhost%s",port)
	log.Fatal(http.ListenAndServe(port,  nil))
}
