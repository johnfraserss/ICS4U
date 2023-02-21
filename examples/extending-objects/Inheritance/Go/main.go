package main

import "fmt"

func main() {
	// Create an object representing Marge (Person)
	marge := Person{
		firstName: "Marge",
		lastName:  "Simpson",
		age:       36,
	}

	// Create an object representing Homer (Employee)
	homer := Employee{
		// Specifying details for parent class
		Person: Person{
			firstName: "Homer",
			lastName:  "Simpson",
			age:       28,
		},
		staffNumber: 1007,
	}

	fmt.Println(marge)
	fmt.Println(homer)
}
