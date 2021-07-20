package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/hoonn9/nomadcoin/blockchain"
)

const port string = ":4000"

// template render 도 private public 영향을 받음
type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}
func home(rw http.ResponseWriter, r *http.Request) {

	// Fprint => writer 에 formatting 해서 출력
	// fmt.Fprint(rw, "Hello from home!")

	// template Must => error 출력 후 panic 발생 해줌
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	// route
	http.HandleFunc("/", home)

	// go server
	fmt.Printf("Listening on http://localhost%s\n", port)
	// log.Fatal 에러가 있으면 출력 하고 os.Exit(1) 프로그램 종료
	log.Fatal(http.ListenAndServe(port, nil))
}
