/**
 * A Object holding the price, author, and title of a book
 * 
 * @author Mr. Seidel (modified from https://thimble.mozilla.org/en-US/anonymous/1dbb9564-0bc8-4a2e-94e8-ba6f6690fb69/304196)
 * @since JDK 1.8
 * @version 1.0
 *
 * Updated 13-Apr-20:
 *  - "printX" methods
 *  + "getX" methods
 *  + updated main() function
 */
class Book {

  private String author;
  private String title;
  private double price;

  /**
   * This is the constructor function for new books
   * 
   * @param author - This is the initial author
   * @param title - This is the initial title of the book
   *
   */
  public Book (String author, String title){
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
  public Book (String author, String title, double price){
    this.price = price;
    this.author = author;
    this.title = title;
  }

  /**
   * Returns the author of the book
   *
   * @returns the author of the book as a String
   */
  public String getAuthor() {
    return this.author;
  }

  /**
   * Returns the title of the book
   *
   * @returns the title of the book as a String
   */
  public String getTitle() {
    return this.title;
  }

  /**
   * Returns the price of the book
   *
   * @returns the price of the book as a double
   */
  public double getPrice() {
    return this.price;
  }

  /**
   * Attempts to increase the price of the book
   * 
   * @param increase This is the amount to increase the price of the book
   * @return true if it is successful in modifying the price, false otherwise.
   */
  public boolean increasePrice(double increase) {
    //Can put in a section here to ensure success of increase.
    //Going to assume success in adding the increase to the price.

    price = price + increase;
    return true;
  }
}

/*

Note for those using Processing.  You don't HAVE to put in the words "public", as all
methods and variables without this prefix are ASSUMED to be "public"

*/


class Main {
  public static void main(String[] args) {
    Book bookOne = new Book("Terry Pratchett", "Guards! Guards!", 5.99);
    Book bookTwo = new Book("Robert Jordan", "The Eye of the World", 8.99);
    System.out.println(bookOne.getAuthor());
    System.out.println(bookOne.getPrice());
    System.out.println(bookOne.getTitle());
    bookOne.increasePrice(1.00);
    System.out.println(bookOne.getPrice());

    System.out.println(bookTwo.getAuthor());
    System.out.println(bookTwo.getPrice());
    System.out.println(bookTwo.getTitle());
    bookTwo.increasePrice(-6.00);
    System.out.println(bookTwo.getPrice());
  }
}