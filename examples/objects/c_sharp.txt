/// <summary>
/// A Object holding the price, author, and title of a book
/// </summary>
/// <remarks>
/// Modified example from https://thimble.mozilla.org/en-US/anonymous/1dbb9564-0bc8-4a2e-94e8-ba6f6690fb69/304196
/// </remarks>
public class Book {

  /// <summary>
  /// Stores the author property as a String
  /// </summary>
  private String author;

  /// <summary>
  /// Stores the title property as a String
  /// </summary>
  private String title;

  /// <summary>
  /// Stores the price property as a float
  /// </summary>
  private float price;

  /// <summary>
  /// The class constructor
  /// </summary>
  /// <param name="author"> Initial author of the book</param>
  /// <param name="title"> Initial title of the book</param>
  public Book (String author, String title){
    this.price = 5.00;
    this.author = author;
    this.title = title;
  }

  /**
   * Prints out the author of the book
   */
  void printAuthor() {
    Console.Writeline(author);
  }

  /// <summary>
  /// Prints out the title of the book
  /// </summary>
  void printTitle() {
    Console.Writeline(title);
  }

  /// <summary>
  /// Prints out the price of the book
  /// </summary>
  void printPrice() {
    Console.Writeline(price);
  }

  /// <summary>
  /// Attempts to increase the price of the book
  /// </summary>
  /// <param="increase"> Price to increase by </param>
  /// <returns> true if it is successful in modifying the price, false otherwise. </returns>
  boolean increasePrice(float increase) {
    //Can put in a section here to ensure success of increase.
    //Going to assume success in adding the increase to the price.

    price = price + increase;
    return true;
  }
}
