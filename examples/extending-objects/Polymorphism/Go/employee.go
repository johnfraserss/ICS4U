package main

import "fmt"

type Employee struct {
	Person
	staffNumber uint
}

// Method function on Employee, returns a string representation of the employee
func (employee Employee) String() string {
	return fmt.Sprintf("%s, %d", employee.Person, employee.staffNumber)
}
