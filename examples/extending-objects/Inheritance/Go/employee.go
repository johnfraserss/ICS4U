package main

import "fmt"

type Employee struct {
	Person
	staffNumber uint
}

func (employee Employee) String() string {
	return fmt.Sprintf("%s, %d", employee.Person, employee.staffNumber)
}
