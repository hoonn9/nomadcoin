package main

import "fmt"

func main() {
	// array
	// go는 array 크기를 정해줘야 한다.
	// array는 append 불가
	// foods := [3]string{"potato", "pizza", "pasta"}

	// for _, foods := range foods {
	// 	fmt.Println(foods)
	// }

	// for i := 0; i<len(foods); i++ {
	// 	fmt.Println(foods[i])
	// }

	// slice 배열은 무한히 커질 수 있다.
	// go가 알아서 공간 늘려줌
	foods := []string{"potato", "pizza", "pasta"}

	fmt.Printf("%v\n", foods)
	// append 는 element 를 추가한 배열을 반환
	foods = append(foods, "tomato")
	fmt.Printf("%v\n", foods)

}
