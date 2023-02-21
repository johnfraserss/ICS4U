package main

import "fmt"

// Inherits a Person object, and adds on a staffNumber
type Employee struct {
	Person
	staffNumber uint
}

// Returns a string representation of the employee
func (employee Employee) String() string {
	return fmt.Sprintf("%s, %d", employee.Person, employee.staffNumber)
}
