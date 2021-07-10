package main

import (
	"fmt"

	"github.com/hoonn9/nomadcoin/person"
)



func main() {
	// export는 Uppercase 시작으로 설정
	// struct field도 export 하면 직접 접근이 불가능하다.
	hoon := person.Person{}
	hoon.SetDetails("hoon", 26)
	fmt.Println("Main 'hoon'", hoon)
}
