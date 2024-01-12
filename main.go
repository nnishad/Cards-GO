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

	//jimPointer := &jim
	jim.updateName("CoolBoy")

	jim.print()
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Println()
	fmt.Printf("%+v", p)
}
