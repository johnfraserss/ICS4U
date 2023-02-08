package main

import "fmt"

type Human interface {
	String() string
}

type Person struct {
	firstName, lastName string
	age                 uint
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s, %d", p.firstName, p.lastName, p.age)
}
