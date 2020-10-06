const Book = require('./book');

/**
 * Function used to test the creation of objects
 */
function testing(){
    var bookOne = new Book("Terry Pratchett", "Guards! Guards!", 5.99);
    var bookTwo = new Book("Robert Jordan", "The Eye of the World", 8.99);
    console.log(bookOne.getAuthor());
    console.log(bookOne.getPrice());
    console.log(bookOne.getTitle());
    bookOne.increasePrice(1.00);
    console.log(bookOne.getPrice());

    console.log(bookTwo.getAuthor());
    console.log(bookTwo.getPrice());
    console.log(bookTwo.getTitle());
    bookTwo.increasePrice(-6.00);
    console.log(bookTwo.getPrice());
}

testing();
