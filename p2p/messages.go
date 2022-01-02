package p2p

import (
	"encoding/json"

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

func (m *Message) addPayload(p interface{}) {
	b, err := json.Marshal(p)
	utils.HandleErr(err)

	m.Payload = b
}

func makeMessage(kind MessageKind, payload interface{}) []byte {
	m := Message {
		Kind: kind,
	}

	m.addPayload(payload)

	mJson, err := json.Marshal(m)
	utils.HandleErr(err)
	return mJson
}

func sendNewestBlock(p *peer) {
	b, err := blockchain.FindBlock(blockchain.Blockchain().NewestHash)
	utils.HandleErr(err)

	m := makeMessage(MessageNewestBlock, b)
	p.inbox <- m
}