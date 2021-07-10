package main

import "fmt"

type person struct {
	name string
	age int
}

// struct method 생성
// (<struct variable name> <struct>)<function name>(arguments) <returnType> {}
// 흔히 struct 명의 첫 자를 변수명으로 사용함
func (p person) sayHello() {
	fmt.Printf("Hello! My name is %s and I'm %d", p.name, p.age)
}



func main() {
	/*
		go는 class, object가 없다.
		struct가 그 역할을 함.
	*/
	// hoon := person{"hoon", 26}
	hoon := person{name: "hoon",age: 26}
	hoon.sayHello()
}
