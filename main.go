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
	//alex := person{firstName: "Alex", lastName: "Anderson"}
	//fmt.Println(alex)
	//
	var alex person
	alex.firstName = "Alex"

	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "abc@gmail.com",
			zipCode: 98989,
		},
	}

	jim.print()

	jim.updateName("CoolBoy")

	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)
}
