package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (n person) speak() {
	fmt.Println("Name:", n.firstName, n.lastName, "/n age:", n.age)
}
