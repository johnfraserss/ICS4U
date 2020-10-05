namespace bookExample {
	class MainClass {
	  public static void Main (string[] args) {
		Book bookOne = new Book("Terry Pratchett", "Guards! Guards!", 5.99);
		Book bookTwo = new Book("Robert Jordan", "The Eye of the World", 8.99);
		Console.WriteLine(bookOne.getAuthor());
		Console.WriteLine(bookOne.getPrice());
		Console.WriteLine(bookOne.getTitle());
		bookOne.increasePrice(1.00);
		Console.WriteLine(bookOne.getPrice());

		Console.WriteLine(bookTwo.getAuthor());
		Console.WriteLine(bookTwo.getPrice());
		Console.WriteLine(bookTwo.getTitle());
		bookTwo.increasePrice(-6.00);
		Console.WriteLine(bookTwo.getPrice());
	  }
	}
}