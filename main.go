package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Welcome to Nomad Coin\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer:	Start the HTML Explorer\n")
	fmt.Printf("rest:	Start the REST API(recommanded)\n")
	// error code 0 => 에러 없음
	os.Exit(0)
}

// CLI (Command Line Interface)
// only using flag package (cobra's famous cli framework but not using)

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	// command
	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
	default:
		usage()
	}
}
