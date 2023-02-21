package main

// A struct representing a book, storing information about it
// Note that Go doesn't actually have classes (it's not really OOP https://go.dev/doc/faq#Is_Go_an_object-oriented_language),
// hence why structs are used instead.
type Book struct {
	Author, Title string
	Price         float64
}

// Constructor function to create a new book.
//
// Parameters:
//
//	author, string:
//	  the author of the book
//	title, string:
//	  the title of the book
//	price, float64:
//	  the decimal price of the book
//
// Returns a pointer to new Book object
func NewBook(author string, title string, price float64) *Book {
	book := new(Book)
	book.Author = author
	book.Title = title
	book.Price = price

	return book
}

// Method on the struct Book that attempts to increase the price of a book.
//
// Parameters:
//
//	increase, float64:
//	   the decimal value to try and increase the price of the book by
//
// Returns whether it was successful in modifying the price
func (book *Book) IncreasePrice(increase float64) bool {
	if increase < 0 {
		// This is a decrease, not an increase
		return false
	}
	book.Price += increase
	return true
}
