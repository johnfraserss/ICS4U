package main

import "fmt"

func main() {
	marge := AbstractPerson{
		firstName: "Marge",
		lastName:  "Simpson",
		age:       36,
	}

	homer := Employee{
		Person: AbstractPerson{
			firstName: "Homer",
			lastName:  "Simpson",
			age:       28,
		},
		staffNumber: 1007,
	}

	simpsons := make([]Person, 2)
	simpsons[0] = marge
	simpsons[1] = homer

	for i := 0; i < len(simpsons); i++ {
		fmt.Println(simpsons[i])
	}
}
