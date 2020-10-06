/**
 * This is the class used to create Book objects
 * This has been updated to use ES6 script
 * 
 * @class
 */
class Book {

  /**
   * This function is the constructor function for a Book object (modified from https://thimble.mozilla.org/en-US/anonymous/1dbb9564-0bc8-4a2e-94e8-ba6f6690fb69/304196)
   * This has been updated to use ES6 script
   * @constructor
   * @param {string} author - This is the initial author (will be a private instance variable)
   * @param {string} title - This is the initial title of the book (will be a private instance variable)
   * @param {number} price - This is the initial price of the book (defaults to 0.00 if no value given)
   */
  constructor(author, title, price=0.00) {
    this.author = author;
    this.title = title;
    this.price = price;
  }

  /**
   * Returns the author of the book
   *
   * @returns {string} - the author of the book
   */
  getAuthor () {
    return this.author;
  }

  /**
   * Returns the title of the book
   *
   * @returns {string} - the title of the book
   */
  getTitle () {
    return this.title;
  }

  /**
   * Returns the price of the book
   *
   * @returns {number} - the price of the book
   */
  getPrice () {
    return this.price;
  }


  /**
   * Increases the price of the book based on the value given
   * This has been updated to use ES6 script
   *  
   * @param {number} value - This is the value to increase the price by.
   * @returns {bool} - true if this succeeded, false if the value would decrease
   */
  increasePrice (value) {
    if (value < 0) {
      return false
    }
    this.price = this.price + value;
    return true;
  }
}

// export the book class
module.exports = Book;