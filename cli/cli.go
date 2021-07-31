package cli

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/hoonn9/nomadcoin/explorer"
	"github.com/hoonn9/nomadcoin/rest"
	"github.com/hoonn9/nomadcoin/utils"
)

func usage() {
	fmt.Printf("Welcome to Nomad Coin\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n")
	// error code 0 => 에러 없음
	// os.Exit(0)

	// defer 는 실행하고 종료
	runtime.Goexit()
}

// CLI (Command Line Interface)
// only using flag package (cobra's famous cli framework but not using)
func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.String("port", "4000", "Set port of th server. if mode is all, input two port ex) 3000:4000")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest' and 'all'")

	flag.Parse()

	switch *mode {
	case "rest":
		port, err := strconv.Atoi(*port)
		utils.HandleErr(err)
		rest.Start(port)
	case "html":
		port, err := strconv.Atoi(*port)
		utils.HandleErr(err)
		explorer.Start(port)
	case "all":
		slices := strings.Split(*port,":")

		if len(slices) < 2 {
			log.Fatal(errors.New("all mode required two port. <rest port>:<html port> (ex: 4000:5000)"))
		}

		for index, item := range(slices) {
			port, err := strconv.Atoi(item)
			utils.HandleErr(err)
			if index == 0 {
				go rest.Start(port)
			} else {
				explorer.Start(port)
			}
		}
	default:
		usage()
	}
}