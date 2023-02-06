package main

import "fmt"

func main() {
	book1 := Book{"George Orwell", "Nineteen Eighty-Four", 10.99}
	book2 := Book{"Robert Jordan", "The Eye of the World", 8.99}

	fmt.Println("Book 1:")
	fmt.Printf("  Title: %s\n", book1.Title)
	fmt.Printf("  Author: %s\n", book1.Author)
	fmt.Printf("  Price: $%.2f\n", book1.Price)
	fmt.Println()

	fmt.Println("Book 2:")
	fmt.Printf("  Title: %s\n", book2.Title)
	fmt.Printf("  Author: %s\n", book2.Author)
	fmt.Printf("  Price: $%.2f\n", book2.Price)
	fmt.Println()

	fmt.Println("Book 2 price increase success:", book2.IncreasePrice(15))
	fmt.Println()

	fmt.Println("Book 2 (after price increase):")
	fmt.Printf("  Title: %s\n", book2.Title)
	fmt.Printf("  Author: %s\n", book2.Author)
	fmt.Printf("  Price: $%.2f\n", book2.Price)
	fmt.Println()

	fmt.Println("Book 1 price increase success:", book1.IncreasePrice(-12))
	fmt.Println()

	fmt.Println("Book 1 (after attempted price increase):")
	fmt.Printf("  Title: %s\n", book1.Title)
	fmt.Printf("  Author: %s\n", book1.Author)
	fmt.Printf("  Price: $%.2f\n", book1.Price)
	fmt.Println()
}
