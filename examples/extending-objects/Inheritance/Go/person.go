package main

import "fmt"

// A struct representing a Person, storing information about it
// Note that Go doesn't actually have classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why structs are used instead.
type Person struct {
	firstName, lastName string
	age                 uint
}

// Method function on Person that performs string conversion.
//
// Returns a string representation of an Person.
func (person Person) String() string {
	return fmt.Sprintf("%s %s, %d", person.firstName, person.lastName, person.age)
}
