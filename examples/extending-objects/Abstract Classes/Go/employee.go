package main

import "fmt"

type Employee struct {
	Person
	staffNumber uint
}

func (emp Employee) String() string {
	return fmt.Sprintf("%s, %d", emp.Person, emp.staffNumber)
}
