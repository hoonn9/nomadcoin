package main

import "fmt"

// func plus (a, b int, name string) (int, string) {
// 	return a + b, name
// }

// func plus (a ...int) (int) {
// 	total := 0
// 	for _, item := range a {
// 		total += item
// 	}
// 	return total
// }



func main() {
	// result := plus(2, 3, 4, 5, 6, 7)
	name := "hoon ! ! ! ! ! ! Is my name"

	// 기본적으로 byte로 넘겨줌
	for _, letter := range name {
		fmt.Println(string(letter))
		fmt.Printf("%x", letter)
	}
}
