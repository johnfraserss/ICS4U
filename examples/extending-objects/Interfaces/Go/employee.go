package main

import "fmt"

// A struct representing an Employee
// Note that Go doesn't actually have classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why structs are used instead.
type Employee struct {
    firstName, lastName string
    age, staffNumber    uint
}

// Method function on Employee that performs string conversion.
//
// Returns a string representation of an Employee.
func (employee Employee) String() string {
    return fmt.Sprintf("%s %s, %d, %d", employee.firstName, employee.lastName, employee.age, employee.staffNumber)
}
