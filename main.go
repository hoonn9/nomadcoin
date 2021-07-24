package main

import (
	"github.com/hoonn9/nomadcoin/explorer"
	"github.com/hoonn9/nomadcoin/rest"
)

func main() {
	// rest 실행 안됨
	// go routine 실행
	// router 중복 에러 => http router 모듈이 서로 같은 것을 처리하고 있음
	// 같은 multiplexer 를 사용하고 있는 문제
	go explorer.Start(3000)
	rest.Start(4000)

	/*
		Multiplexer => url 을 보고 request를 처리하고 handler 호출
	*/
}
