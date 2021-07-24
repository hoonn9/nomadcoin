package main

import (
	"flag"
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
// command  ex) start
// flag  ex) -port=3000  -https
func main() {
	// fmt.Println(os.Args[2:])
	if len(os.Args) < 2 {
		usage()
	}
	// flag set  => command 에 따른 flag 들을 set
	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	// 타입 설정, flag 명, default 값, 선택 시 띄워줄 문구
	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	// command
	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
		fmt.Println("Start REST API")
	default:
		usage()
	}

	// rest가 parse 됐을 때
	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start server")
	}
}
