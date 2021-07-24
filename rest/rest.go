package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hoonn9/nomadcoin/blockchain"
	"github.com/hoonn9/nomadcoin/utils"
)

var port string

type url string

// Json 을 반환할 때 가공해서 반환 MarshalText
func (u url) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("http://localhost%s%s", port, u)), nil
}

// type URLDescription struct implements ... 대신 interface method
type urlDescription struct {
	URL 		url 	`json:"url"`
	Method 		string 	`json:"method"`
	Description string 	`json:"description"`
	Payload		string 	`json:"payload,omitempty"`
}

type addBlockBody struct {
	Message string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL: url("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL: url("/blocks"),
			Method: "POST",
			Description: "Add A Block",
			Payload: "data:string",
		},
		{
			URL: url("/blocks/{id}"),
			Method: "GET",
			Description: "See A Block",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}

func Start(aPort int) {
	// Mux === MultiPlexer
	handler := http.NewServeMux()

	port = fmt.Sprintf(":%d", aPort)

	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)

	fmt.Printf("Listening on http://localhost%s\n",port)
	// handler nil이면 default mux
	log.Fatal(http.ListenAndServe(port,  handler))
}