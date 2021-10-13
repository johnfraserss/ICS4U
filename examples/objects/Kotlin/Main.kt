// Assuming everything is in the same package

fun main(args : Array<String>) {
    val bookOne = Book("Terry Pratchett", "Guards! Guards!", 5.99)
    val bookTwo = Book("Robert Jordan", "The Eye of the World", 8.99)
    
    println(bookOne.getAuthor());
    println(bookOne.getPrice());
    println(bookOne.getTitle());
    bookOne.increasePrice(1.00);
    println(bookOne.getPrice());

    println(bookTwo.getAuthor());
    println(bookTwo.getPrice());
    println(bookTwo.getTitle());
    bookTwo.increasePrice(-6.00);
    println(bookTwo.getPrice());
}