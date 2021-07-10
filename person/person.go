package person

type Person struct {
	name string
	age int
}

// receiver function
// 포인터를 안쓰면 복사본 => 원본을 updating 하려면 포인터
// func (p Person) SetDetails(name string, age int) {
// 	p.name = name
// 	p.age = age
// 	fmt.Println("See detail hoon:", p)
// }

func (p *Person) SetDetails(name string, age int) {
	p.name = name
	p.age = age
}

func (p Person) Name() string {
	return p.name
}