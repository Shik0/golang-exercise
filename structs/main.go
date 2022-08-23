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
	alex := person{
		firstName: "Alex",
		lastName:  "John",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 90000,
		},
	}

	alex.updateName("Alexander")
	alex.print()
}

func (personNamePointer *person) updateName(newFirstname string) {
	(*personNamePointer).firstName = newFirstname
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
