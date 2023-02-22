package main

import "fmt"

// Defines what functions an object that we want to treat as a Human should have
type Human interface {
	String() string
}

// A struct representing an Person (base class)
// Note that Go doesn't actually have classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why structs are used instead.
type Person struct {
	firstName, lastName string
	age                 uint
}

// Method function on Person that performs string conversion.
//
// Returns a string representation of an Person.
func (p Person) String() string {
	return fmt.Sprintf("%s %s, %d", p.firstName, p.lastName, p.age)
}
