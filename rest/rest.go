package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

type errResponse struct {
	ErrorMessage string `json:"errorMessage"`
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
			URL: url("/blocks/{height}"),
			Method: "GET",
			Description: "See A Block",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlocks())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	default:
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func block(rw http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	// String to Integer Atoi
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)

	block, err := blockchain.GetBlockchain().GetBlock(id)

	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		// Sprint 는 DefaultFormat 사용
		encoder.Encode(errResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}
}

// middleware
func jsonContentTypeMiddleware(next http.Handler) http.Handler {

	// HandleFunc 은 type이다.
	// 이 함수의 return 타입인 http.Handler 에 오류를 안띄우는 이유는
	// HandleFunc의 receiver method 가 arguments에 따라 동적인 return 을 주기 때문이다. (adaptor 패턴)
	
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	router := mux.NewRouter()
	port = fmt.Sprintf(":%d", aPort)
	
	// Middleware
	router.Use(jsonContentTypeMiddleware)

	// Methods => request method 제한
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")

	fmt.Printf("Listening on http://localhost%s\n",port)
	// handler nil이면 default mux
	log.Fatal(http.ListenAndServe(port,  router))
}