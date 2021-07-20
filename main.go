package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/hoonn9/nomadcoin/blockchain"
)

const (
	port string = ":4000"
	templateDir string = "templates/"
)
var templates *template.Template

// template render 도 private public 영향을 받음
type homeData struct {
	PageTitle string
	Blocks []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	templates.ExecuteTemplate(rw, "home", data)
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func main() {
	// go는 ** 지원 안함
	// template pattern load
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))

	// route
	http.HandleFunc("/", home)
	http.HandleFunc("/add", add)

	// go server
	fmt.Printf("Listening on http://localhost%s\n", port)
	// log.Fatal 에러가 있으면 출력 하고 os.Exit(1) 프로그램 종료
	log.Fatal(http.ListenAndServe(port, nil))
}
