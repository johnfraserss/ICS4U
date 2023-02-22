package main

import "fmt"

// A struct representing an employee using the abstract class Person, storing information about it
// Note that Go doesn't actually have classes, or abstract classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why structs are used instead.
type Employee struct {
	Person
	staffNumber uint
}

// Method function on Employee that performs string conversion.
//
// Returns a string representation of an Employee.
func (emp Employee) String() string {
	return fmt.Sprintf("%s, %d", emp.Person, emp.staffNumber)
}
