
import Foundation

let bookOne = Book(author: "Terry Pratchett", title: "Guards! Guards!", price: 5.99)
let bookTwo = Book(author: "Robert Jordan", title: "The Eye of the World", price: 8.99)

print(bookOne.getAuthor())
print(bookOne.getPrice())
print(bookOne.getTitle())
bookOne.increasePrice(1.00)
print(bookOne.getPrice())

print(bookTwo.getAuthor())
print(bookTwo.getPrice())
print(bookTwo.getTitle())
bookTwo.increasePrice(-6.00)
print(bookTwo.getPrice())






