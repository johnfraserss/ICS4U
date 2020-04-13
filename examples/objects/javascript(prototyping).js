/**
 * This is the class used to create Book objects
 * This is using the Prototyping methodology
 * This function is both the original class definition and the constructor at the same time 
 *
 * @class
 * @constructor
 * @param {string} author - This is the initial author (will be a private instance variable)
 * @param {string} title - This is the initial title of the book (will be a private instance variable)
*/
let Book = function(author, title) {
  this.author = author;
  this.title = title;
  this.price = 0.00;
}

let Book = function(author, title, price) {
  this.author = author;
  this.title = title;
  this.price = price;
}

/**
* Returns the author of the book
*
* @returns {string}
*/
Book.prototype.getAuthor = function() {
  return this.author;
}

/**
* Returns the title of the book
*
* @returns {string}
*/
Book.prototype.getTitle = function() {
  return this.title;
}

/**
* Returns the price of the book
*
* @returns {number}
*/
Book.prototype.getPrice = function() {
  return this.price;
}

/**
 * Increases the price of the book based on the value given
 * This has been updated to use ES6 script
 *  
 * @param {number} value - This is the value to increase the price by.
 */
Book.prototype.increasePrice = function (value) {
  this.price = this.price + value;
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
