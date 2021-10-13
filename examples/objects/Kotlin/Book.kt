


/**
 * An object holding the price, author, and title of a book
 * 
 * @author Mr. Seidel (modified from https://thimble.mozilla.org/en-US/anonymous/1dbb9564-0bc8-4a2e-94e8-ba6f6690fb69/304196)
 *
 * Updated 13-Oct-21:
 * converted from the Java version
 * 
 * @property author - This is the initial author
 * @property title - This is the initial title of the book
 * @property price - This is the initial price of the book
 *
 */
class Book (_author: String, _title: String, _price: Double) {

  private var author: String;
  private var title: String;
  private var price: Double;

  /**
   * @constructor
   */
  init {
      author = _author
      title = _title
      price = _price
  }
    
  /**
   * Returns the author of the book
   *
   * @returns the author of the book as a String
   */
  fun getAuthor(): String {
    return author;
  }

  /**
   * Returns the title of the book
   *
   * @returns the title of the book as a String
   */
  fun getTitle(): String {
    return title;
  }

  /**
   * Returns the price of the book
   *
   * @returns the price of the book as a double
   */
  fun getPrice(): Double {
    return price;
  }

  /**
   * Attempts to increase the price of the book
   * 
   * @param increase This is the amount to increase the price of the book
   * @return true if it is successful in modifying the price, false otherwise.
   */
  fun increasePrice(increase: Double): Boolean {
    if (increase < 0) { // this is a decrease, not increase
      return false;
    }
    price = price + increase;
    return true;
  }
}
