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
*
* @return the author of the book
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
*
* @return the price of the book
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
	if (increase < 0) {	// this is a decrease, not an increase
		return false;
	}
	price = price + increase;
	return true;
}