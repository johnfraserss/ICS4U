package main

import "fmt"

func main() {
    // Initialize some books
    book1 := Book{"George Orwell", "Nineteen Eighty-Four", 10.99}
    book2 := Book{"Robert Jordan", "The Eye of the World", 8.99}

    // Print out info for book 1
    fmt.Println("Book 1:")
    fmt.Printf("  Title: %s\n", book1.Title)
    fmt.Printf("  Author: %s\n", book1.Author)
    fmt.Printf("  Price: $%.2f\n", book1.Price)
    fmt.Println()

    // Print out info for book 2
    fmt.Println("Book 2:")
    fmt.Printf("  Title: %s\n", book2.Title)
    fmt.Printf("  Author: %s\n", book2.Author)
    fmt.Printf("  Price: $%.2f\n", book2.Price)
    fmt.Println()

    // Increase price of book 2
    fmt.Println("Book 2 price increase success:", book2.IncreasePrice(15))
    fmt.Println()

    // Print out info for book 2 after increase
    fmt.Println("Book 2 (after price increase):")
    fmt.Printf("  Title: %s\n", book2.Title)
    fmt.Printf("  Author: %s\n", book2.Author)
    fmt.Printf("  Price: $%.2f\n", book2.Price)
    fmt.Println()

    // Try to increase price of book 1
    fmt.Println("Book 1 price increase success:", book1.IncreasePrice(-12))
    fmt.Println()

    // Print out info for book 1 after attempted increase
    fmt.Println("Book 1 (after attempted price increase):")
    fmt.Printf("  Title: %s\n", book1.Title)
    fmt.Printf("  Author: %s\n", book1.Author)
    fmt.Printf("  Price: $%.2f\n", book1.Price)
    fmt.Println()
}
