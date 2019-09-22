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
   */
  constructor(author, title) {
    this.author = author;
    this.title = title;
  }

  /**
   * Prints out the author of the book
   */
  printAuthor () {
    console.log(this.author);
  }

  /**
   * Prints out the author of the book
   */
  printTitle () {
    console.log(this.title);
  }

  /**
   * Prints out the author of the book
   */
  printPrice () {
    console.log(this.price);
  }


  /**
   * Increases the price of the book based on the value given
   * This has been updated to use ES6 script
   *  
   * @param {number} value - This is the value to increase the price by.
   */
  increasePrice (value) {
    this.price = this.price + value;
  }

}

/**
 * Function used to test the creation of objects
 */
function testing(){
  var a = new Book("Terry Pratchett", "Guards! Guards!");
  var b = new Book("Robert Jordan", "The Eye of the World");
  a.printAuthor();
  a.printPrice();
  a.increasePrice(3.33);
  a.increasePrice(2.33);
  a.increasePrice(4.33);
  a.printPrice();
  
  b.printAuthor();
  b.printPrice();
  b.increasePrice(1);
  b.printPrice();
}

testing();
