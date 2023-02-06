// A struct representing a book, storing information about it
// Note that Go doesn't actually have classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why structs are used instead.
package main

type Book struct {
	Author, Title string
	Price         float64
}

// IncreasePrice attempts to increase the price of a book.
// Returns whether it was successful in modifying the price.
func (book *Book) IncreasePrice(increase float64) bool {
	if increase < 0 {
		// This is a decrease, not an increase
		return false
	}
	book.Price += increase
	return true
}
