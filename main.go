package main

import (
	"fmt"
	"time"
)

// chan<- => only send channel
func countToTen(c chan<- int)  {
	// 10번 blocking
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d", i)
		c <- i
		
 	}
	 // go routine close
	 close(c)
}

// <-chan => only receive channel
func receive(c <-chan int) {

	for {
		// ok: boolean => channel open 여부
		a, ok := <- c

		if !ok {
			fmt.Println("done.")
			break
		}
		fmt.Printf("received %d\n", a)
	}
}

func main() {

	// blocking operation : 프로그램을 기다림 (a := <- c, for {})
	// 끝난 go routine의 channel message를 계속 기다리면 deadlock

	c := make(chan int)

	go countToTen(c)
	receive(c)
}