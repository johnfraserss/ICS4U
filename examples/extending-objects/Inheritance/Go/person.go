package main

import "fmt"

type Person struct {
	firstName, lastName string
	age                 uint
}

// Method function on Employee, returns a string representation of a Person
func (person Person) String() string {
	return fmt.Sprintf("%s %s, %d", person.firstName, person.lastName, person.age)
}
