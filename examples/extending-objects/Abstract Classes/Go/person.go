// Golang doesn't really have Abstract Classes, so this is an alternative way of implementing them
package main

import "fmt"

// A interface representing the abstract interface of Person
// Note that Go doesn't actually have classes or abstract classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why interfaces are used instead.
type Person interface {
    String() string
    IncreaseAge()
}

// A struct representing the abstract type Person. Includes data about the person as well.
// Note that Go doesn't actually have classes or abstract classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why structs are used instead.
type AbstractPerson struct {
    firstName, lastName string
    age                 uint
    Person
}

// Method function on AbstractPerson that performs string conversion.
//
// Returns a string representation of an AbstractPerson.
func (p AbstractPerson) String() string {
    return fmt.Sprintf("%s %s, %d", p.firstName, p.lastName, p.age)
}

// Increase Age is not defined as this is the abstract function
