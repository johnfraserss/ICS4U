package main

import "fmt"

type Employee struct {
	Person
	staffNumber uint
}

// Method function on Employee, returns a string representation of the employee
func (emp Employee) String() string {
	return fmt.Sprintf("%s, %d", emp.Person, emp.staffNumber)
}
