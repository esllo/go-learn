package article

import "fmt"

// person struct
type person struct {
	name string
	age  int
}

// constructor
func newPerson(name string, age int) *person {
	p := person{}
	p.name = name
	p.age = age
	return &p
}

func Article_Struct() {
	p := person{}
	p.name = "Hong Gildong"
	p.age = 26

	fmt.Printf("%#v\n", p)

	p2 := person{"Hong Cheolsoo", 26}

	fmt.Printf("%#v\n", p2)

	p3 := person{name: "Hong Jjanggoo", age: 26}

	fmt.Printf("%#v\n", p3)

	// person pointer
	personPointer := new(person)
	// pointer use "." not "->"
	personPointer.name = "Hong Maenggoo"
	personPointer.age = 26

	fmt.Printf("%#v\n", personPointer)

	personConstructor := newPerson("Name", 26)

	fmt.Printf("%#v\n", personConstructor)
}
