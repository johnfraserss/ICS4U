/*
  Basic Class Example
  
  Reference: http://www.cplusplus.com/doc/tutorial/classes/
*/

#include <iostream>
using namespace std;

/**
 * A Object holding the price, author, and title of a book
 * 
 * @author Mr. Seidel
 * @since 23-Sep-2018
 * @version 1.0
 *
 */
class Book {
  string author, title;  //default to private
  double price;			 //default to private

  public:
	Book(string, string);
	Book(string, string, double);
	void printAuthor(void);
	void printTitle(void);
	void printPrice(void);
	void increasePrice(double);
}

	
/**
* This is the constructor function for new books
* 
* @param author - This is the initial author
* @param title - This is the initial title of the book
*
*/
Book::Book (string author, string title){
	this.price = 0.00;
	this.author = author;
	this.title = title;
}

/**
* This is the constructor function for new books
* 
* @param author - This is the initial author
* @param title - This is the initial title of the book
* @param price - This is the initial price of the book
*
*/
Book::Book (string author, string title, double price){
	this.price = price;
	this.author = author;
	this.title = title;
}

/**
* Prints out the author of the book
*/
void Book::printAuthor() {
	cout << this.author << endl;
}

/**
* Prints out the title of the book
*/
void Book::printTitle() {
	cout << this.title << endl;
}

/**
* Prints out the price of the book
*/
void Book::printPrice() {
	cout << this.price << endl;
}

/**
* Attempts to increase the price of the book
* 
* @param increase This is the amount to increase the price of the book
* @return true if it is successful in modifying the price, false otherwise.
*/
boolean Book::increasePrice(double increase) {
	//Can put in a section here to ensure success of increase.
	//Going to assume success in adding the increase to the price.

	price = price + increase;
	return true;
}
