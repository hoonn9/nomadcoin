package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/hoonn9/nomadcoin/blockchain"
	"github.com/hoonn9/nomadcoin/utils"
)


type MessageKind int

// iota type
// 선언한 아래 부터 순서대로 설정됨 (0,1,2,3,4)
const (
	MessageNewestBlock				MessageKind = iota
	MessageAllBlocksBlock
	MessageAllBlocksResponse

)

type Message struct {
	Kind	MessageKind
	Payload	[]byte
}


func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message {
		Kind: kind,
		Payload: utils.ToJSON(payload),
	}
	return utils.ToJSON(m)
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)

	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}

func handleMsg(m *Message, p *peer) {
	switch m.Kind {
	case MessageNewestBlock:
		var payload blockchain.Block
		utils.HandleErr(json.Unmarshal(m.Payload, &payload))
		fmt.Println(payload)
	}

}