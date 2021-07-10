package main

import "fmt"


func main() {
	x := 31254125234
	// binary 2
	// sprintf => format 된 데이터 반환
	xAsBinary := fmt.Sprintf("%b\n", x)
	fmt.Println(x, xAsBinary)
	// // 8
	// fmt.Printf("%o\n", x)
	// // 16
	// fmt.Printf("%x\n", x)


}
