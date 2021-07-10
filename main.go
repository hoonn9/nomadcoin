package main

import "fmt"

func main() {
	/*
		포인터
		빠르게 만들고 data를 mutate 하는걸 간단하게 해준다.
	*/

	/*
	a := 2
	// a를 copy
	b := a
	a = 12
	fmt.Println(b)
	// 결과: 2
	fmt.Println(&b, &a)
	*/

	a := 2
	// a의 메모리 주소를 복사
	b := &a
	a = 50
	// b에 복사된 메모리 주소에 있는 값: *b
	fmt.Println(*b)
	// 결과: 50
	fmt.Println(b, &a)
}
