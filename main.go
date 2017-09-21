package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println(alex)

	var alex1 person
	alex1.firstName = "Alex"
	alex1.lastName = "Anderson"
	fmt.Println(alex1)
}
