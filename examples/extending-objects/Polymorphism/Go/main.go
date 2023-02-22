package main

import "fmt"

func main() {
	// Make an array of Humans, which includes both Person and Employee classes
	arr := make([]Human, 2)

	// Add a Person to the array
	arr[0] = Person{
		"Marge",
		"Simpson",
		36,
	}

	// Add an Employee to the array
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
