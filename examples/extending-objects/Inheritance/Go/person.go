package main

import "fmt"

type Person struct {
	firstName, lastName string
	age                 uint
}

func (person Person) String() string {
	return fmt.Sprintf("%s %s, %d", person.firstName, person.lastName, person.age)
}
