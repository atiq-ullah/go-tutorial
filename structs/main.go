package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	sultan := person{firstName: "Sultan", lastName: "Ullah", contactInfo: contactInfo{email: "sultan@test.com", zipCode: 9400}}
	sultan.print()
	sultan.updateName("NewName")
	sultan.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(name string) {
	p.firstName = name
}