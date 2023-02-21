package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, staffNumber    uint
}

// Method function on Employee, returns a string representation of the employee
func (employee Employee) String() string {
	return fmt.Sprintf("%s %s, %d, %d", employee.firstName, employee.lastName, employee.age, employee.staffNumber)
}
