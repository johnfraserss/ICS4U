#include <iostream>
#include <string>
#include "book.cpp"
using namespace std;

int main() {
    Book bookOne = Book("Terry Pratchett", "Guards! Guards!", 5.99);
    Book bookTwo = Book("Robert Jordan", "The Eye of the World", 8.99);

    cout << bookOne.getAuthor() << endl;
    cout << bookOne.getPrice() << endl;
    cout << bookOne.getTitle() << endl;
    bookOne.increasePrice(1.00);
    cout << bookOne.getPrice() << endl;

    cout << bookTwo.getAuthor() << endl;
    cout << bookTwo.getPrice() << endl;
    cout << bookTwo.getTitle() << endl;
    bookTwo.increasePrice(-6.00);
    cout << bookTwo.getPrice() << endl;

}
