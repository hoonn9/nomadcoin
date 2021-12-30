package p2p

import (
	"fmt"
	"net/http"
	"time"

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

		if err != nil {
			conn.Close()
			break
		}
		
		fmt.Printf("just got %s\n\n", p)

		time.Sleep((5 * time.Second))
		message := fmt.Sprintf("New Message: %s", p)
		utils.HandleErr(conn.WriteMessage(websocket.TextMessage, []byte(message)))
	}


}