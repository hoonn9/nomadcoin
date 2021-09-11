package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int)  {
	// 10번 blocking
	for i := range [10]int{} {
		time.Sleep(1 * time.Second)
		fmt.Printf("sending %d", i)
		c <- i
 	}
}

func main() {
	c := make(chan int)

	go countToTen(c)

	// blocking operation : 프로그램을 기다림 (a := <- c, for {})

	for {
		a := <- c
		fmt.Printf("received %d\n", a)
	}

	// 끝난 go routine의 channel message를 계속 기다리면 deadlock
}