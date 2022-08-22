package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "John",
		contact: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 90000,
		},
	}

	fmt.Printf("%+v", alex)
}
