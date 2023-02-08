package main

import "fmt"

type PersonInterface interface {
	String() string
	IncreaseAge()
}

type Person struct {
	firstName, lastName string
	age                 uint
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, %d", p.firstName, p.lastName, p.age)
}

// Increase Age is not defined as this is the abstract function
