package p2p

import (
	"fmt"
	"net/http"
	"strings"

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
	openPort := r.URL.Query().Get("openPort")

	// remoteAddr 에는 실제 요청한 서버의 Open port가 업다.
	result := strings.Split(r.RemoteAddr, ":")
	initPeer(conn, result[0], openPort)

}

func AddToPeer(address, port, openPort string)  {
	conn, _, err := websocket.DefaultDialer.Dial(fmt.Sprintf("ws://%s:%s/ws?openPort=%s", address, port, openPort), nil)
	utils.HandleErr(err)

	initPeer(conn, address, port)
}