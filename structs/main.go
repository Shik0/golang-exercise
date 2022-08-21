package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName  string
	}

	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)
}
