
import Foundation

/// A Book Object that holds the price, author, and title of a book
class Book {

  

    var author : String
    var title : String
    var price : Float

    init(author : String, title : String, price : Float=0.00) {
        
        /**
            Constructor to build a book object

            - Parameters:
                - author: The full name of the author of the book
                - title: The full title of the book
                - price: The price of the book in dollars and cents (example format ###.##)

        */

        self.author = author
        self.price = price
        self.title = title

    }

    func getAuthor() -> String {
        
        /**
            Returns the author of the book

            - Returns: The author of the book as a String

        */
                
                
        
        return self.author
    }

    func getPrice() -> Float {
        
        /**
            Returns the price  of the book

            - Returns: The price  of the book as a Float

        */
        
        return self.price
    }

    func getTitle() -> String {
        
        /**
            Returns the title of the book to the console

            - Returns: The title of the book

        */
        
        return self.title
    }

    func increasePrice(_ increase : Float) -> Bool {
        
        /**
            Attempts to increase the price of the book

            - Parameters:
                - increase: The value to increase the price by. This value can be negative; however, it will never lower the value below zero. If this happens, the function will return false
            
            - Returns: true  if the method was successful
            - Returns: false if the method attempted to bring the value below zero
         
         
        

        */

        if (self.price + increase) > 0 {
            self.price += increase
            return true
        }

        return false

    }

}




