package p2p

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/hoonn9/nomadcoin/utils"
)

var upgrader = websocket.Upgrader{}

func Upgrade(rw http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(rw, r, nil)
	utils.HandleErr(err)


	for {
		// blocking operation
		// 메시지가 도착할 때 까지 block
		_, p, err := conn.ReadMessage()
		utils.HandleErr(err)
		fmt.Printf("%s\n\n", p)
	}



}