package p2p


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

