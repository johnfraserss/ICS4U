#include <iostream>
#include <string>
using namespace std;

/**
 * A Object holding the price, author, and title of a book
 * 
 * @author Mr. Seidel
 * @since 23-Sep-2018
 * @version 1.0
 * 
 * Updated 13-Apr-20:
 *  - "printX" methods
 *  + "getX" methods
 *  + updated main() function
 */
class Book {
    private:
        string author, title;  
        double price;		

    public:
	    Book(string, string);
	    Book(string, string, double);
        ~Book();
	    string getAuthor(void);
	    string getTitle(void);
	    double getPrice(void);
	    bool increasePrice(double);
};  // <--- note the semi-colon here

	
/**
* This is the constructor function for new books
* 
* @param author - This is the initial author
* @param title - This is the initial title of the book
*
*/
Book::Book (string _author, string _title){
	price = 0.00;
	author = _author;
	title = _title;
}

/**
* This is the constructor function for new books
* 
* @param _author - This is the initial author
* @param _title - This is the initial title of the book
* @param _price - This is the initial price of the book
*
*/
Book::Book (string _author, string _title, double _price){
	price = _price;
	author = _author;
	title = _title;
}

/**
* Returns the author of the book
*/
string Book::getAuthor() {
	return author;
}

/**
* Returns the title of the book
* 
* @return the title of the book
*/
string Book::getTitle() {
	return title;
}

/**
* Returns the price of the book as a double
*/
double Book::getPrice() {
	return price;
}

/**
* Attempts to increase the price of the book
* 
* @param increase This is the amount to increase the price of the book
* @return true if it is successful in modifying the price, false otherwise.
*/
bool Book::increasePrice(double increase) {
	//Can put in a section here to ensure success of increase.
	//Going to assume success in adding the increase to the price.

	price = price + increase;
	return true;
}

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