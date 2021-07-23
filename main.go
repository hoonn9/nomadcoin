package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

// Json 을 반환할 때 가공해서 반환 MarshalText
func (u URL) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost%s%s", port, u)), nil
}

// type URLDescription struct implements ... 대신 interface method
type URLDescription struct {
	URL 		URL `json:"url"`
	Method 		string `json:"method"`
	Description string `json:"description"`
	Payload		string `json:"payload,omitempty"`
}

// Stringers
// struct 출력 시 return 값 제어
// Go 는 extends, implements가 없어서 interface가 method로 해결한다.
/*
func (u URLDescription) String() string {
	return "Hello I'm the URL Description"
}
*/


func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: URL("/blocks"),
			Method: "POST",
			Description: "Add A Block",
			Payload: "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}



func main() {
	http.HandleFunc("/", documentation)

	fmt.Printf("Listening on http://localhost%s",port)
	log.Fatal(http.ListenAndServe(port,  nil))
}
