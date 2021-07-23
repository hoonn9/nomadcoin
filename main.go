package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"


/*
	field struct tag
	Go에서 export를 위해 대문자로 시작하는 property Key를 
	json에 맞게 소문자로 가공
	json:"<key>"`
	omitempty => field가 비어있으면 숨겨준다.
*/

type URLDescription struct {
	URL 		string `json:"url"`
	Method 		string `json:"method"`
	Description string `json:"description"`
	Payload		string `json:"payload,omitempty"`
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
		{
			URL: "/blocks",
			Method: "POST",
			Description: "Add A Block",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	// #1 
	// slice to json(byte[])
	// b, err := json.Marshal(data)
	// utils.HandleErr(err)

	// fmt.Fprintf(rw, "%s", b)

	// #2
	json.NewEncoder(rw).Encode(data)

	
}

func main() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on http://localhost%s",port)
	log.Fatal(http.ListenAndServe(port,  nil))
}
