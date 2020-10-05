import Book;

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
