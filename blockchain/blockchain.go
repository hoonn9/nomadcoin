package blockchain

type block struct {
	data string
	hash string
	prevHash string
}

type blockchain struct {
	blocks []block
}


// 단방향 함수
// 입력 값으로 얻은 출력값을 출력값을 통해 입력값을 얻을 수 없음

/*
	하나의 hash만 변해도 전체가 변함. 무효화 된다.

	B1
		b1Hash = (data + "")
	B2
		b2Hash = (data + b1Hash)
	B3
		b3Hash = (data + b2Hash)
*/


/*
	singleton application 내 어디서든 단 하나의 instance를 공유하는 패턴
*/

var b *blockchain

func GetBlockchain() *blockchain {
	// only once execute
	if b == nil {
		b = &blockchain{}
	}
	return b
}

