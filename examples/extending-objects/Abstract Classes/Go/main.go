package main

import "fmt"

func main() {
	homer := Employee{
		Person: AbstractPerson{
			firstName: "Homer",
			lastName:  "Simpson",
			age:       28,
		},
		staffNumber: 1007,
	}

	simpsons := make([]Person, 1)
	simpsons[0] = homer

	for i := 0; i < len(simpsons); i++ {
		fmt.Println(simpsons[i])
	}
}
