package main

import "fmt"

func main() {
	marge := Person{
		firstName: "Marge",
		lastName:  "Simpson",
		age:       36,
	}

	homer := Employee{
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
