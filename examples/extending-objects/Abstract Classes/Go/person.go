// Golang doesn't really have Abstract Classes, so this is an alternative way of implementing them
package main

import "fmt"

type Person interface {
	String() string
	IncreaseAge()
}

type AbstractPerson struct {
	firstName, lastName string
	age                 uint
	Person
}

// Method function on AbstractPerson, returns a string representation of the employee
func (p AbstractPerson) String() string {
	return fmt.Sprintf("%s %s, %d", p.firstName, p.lastName, p.age)
}

// Increase Age is not defined as this is the abstract function
