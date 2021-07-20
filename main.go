package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {

	// Fprint => writer 에 formatting 해서 출력
	fmt.Fprint(rw, "Hello from home!")
}

func main() {
	// route
	http.HandleFunc("/", home)

	// go server
	fmt.Printf("Listening on http://localhost%s\n", port)
	// log.Fatal 에러가 있으면 출력
	log.Fatal(http.ListenAndServe(port, nil))
}
