#include <string>
using namespace std;
/**
 * An Object holding the price, author, and title of a book
 * 
 * @author Mr. Seidel
 * @since 23-Sep-2018
 * @version 1.0
 * 
 * Updated 27-Mar-23:
 *  + moved class declaration here
 */
class Book {
    private:
        string author, title;  
        double price;		

    public:
	    Book(string, string);
	    Book(string, string, double);
     	    string getAuthor(void);
	    string getTitle(void);
	    double getPrice(void);
	    bool increasePrice(double);
};  // <--- note the semi-colon here