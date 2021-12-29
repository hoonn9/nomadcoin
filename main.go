package main

import (
	"fmt"
	"time"
)

// chan<- => only send channel
func send(c chan<- int)  {
	// 10번 blocking
	for i := range [10]int{} {
		fmt.Printf(">> sending %d << \n", i)
		// sender blocking (실행된 함수에서 receiving 중 그 함수도 blocking 된다.)
		c <- i
		fmt.Printf("sent %d \n", i)
		
 	}
	 // go routine close
	 close(c)
}

// <-chan => only receive channel
func receive(c <-chan int) {

	for {
		time.Sleep(10 * time.Second)
		// ok: boolean => channel open 여부
		a, ok := <- c

		if !ok {
			fmt.Println("done.")
			break
		}
		fmt.Printf("|| received %d ||\n", a)
	}
}

func main() {

	// blocking operation : 프로그램을 기다림 (a := <- c, for {})
	// 끝난 go routine의 channel message를 계속 기다리면 deadlock

	// unbuffer channel은 큐 공간이 1개인 기본 상태 (all blocking)

	// buffer channel (n 개당 block처리)
	// 큐에 보내기를 n개 채우고, 받기 할떄마다 하나씩 제거됨
	// [1,2,3,4,5] 한번에 보내기 => 10초 기다린 후 1받고 6보내기 큐에 추가...
	c := make(chan int, 10)

	go send(c)
	receive(c)
}