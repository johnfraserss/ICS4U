package main

import "fmt"

func main() {
	arr := make([]Human, 2)

	arr[0] = Person{
		"Marge",
		"Simpson",
		36,
	}

	arr[1] = Employee{
		Person: Person{
			"Homer",
			"Simpson",
			28,
		},
		staffNumber: 1007,
	}

	for _, person := range arr {
		fmt.Println(person)
	}
}
