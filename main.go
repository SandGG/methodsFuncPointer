package main

import "fmt"

type person struct {
	name, lastName string
	age            int
}

func (p *person) NewName(n, ln string) {
	p.name = n
	p.lastName = ln
}

func newAge(p *person, a int) {
	p.age = a
}

func main() {
	var per = person{name: "Maria", lastName: "Ramos", age: 21}
	var perPointer = &person{name: "Juan", lastName: "Quintana", age: 18}

	//Func pointer
	newAge(&per, 16)
	fmt.Println(per)

	//Methods pointer or value
	per.NewName("Marco", "Diaz")
	perPointer.NewName("Abril", "Sanchez")

	fmt.Println(per, perPointer)
}
