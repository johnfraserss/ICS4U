/**
 * A Object holding the price, author, and title of a book
 * 
 * @author Mr. Seidel (modified from https://thimble.mozilla.org/en-US/anonymous/1dbb9564-0bc8-4a2e-94e8-ba6f6690fb69/304196)
 * @since JDK 1.8
 * @version 1.0
 *
 */
class Book {

  private String author;
  private String title;
  private float price;

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
  public Book (String author, String title, float price){
    this.price = price;
    this.author = author;
    this.title = title;
  }

  /**
   * Prints out the author of the book
   */
  public void printAuthor() {
    System.out.println("Author: " + author);
  }

  /**
   * Prints out the title of the book
   */
  public void printTitle() {
    System.out.println("Title: " + title);
  }

  /**
   * Prints out the price of the book
   */
  public void printPrice() {
    System.out.println("Price: $" + float(int(price * 100)) / 100.0);
  }

  /**
   * Attempst to increase the price of the book
   * 
   * @param increase This is the amount to increase the price of the book
   * @return true if it is successful in modifying the price, false otherwise.
   */
  public boolean increasePrice(float increase) {
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
