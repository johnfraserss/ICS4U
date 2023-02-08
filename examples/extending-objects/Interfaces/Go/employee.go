package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, staffNumber    uint
}

func (employee Employee) String() string {
	return fmt.Sprintf("%s %s, %d, %d", employee.firstName, employee.lastName, employee.age, employee.staffNumber)
}
