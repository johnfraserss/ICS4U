/*
In ECMAScript, abstract classes haven't been implemented like they have been in other OOP-based languages (Python, Java, etc.)
If you want to make an abstract class in Javascript, follow something similar to the pattern described below.
*/
class Person {
    constructor() {
        // using new.target from ES2015, we prevent the Abstract class of Person from directly being instantiated
        if (new.target === Person) {
            throw new TypeError("Can't directly instantiate an Abstract Class");
        }

        // Check the subclass' prototype to see if whoAmI has been newly implemented
        if (this.whoAmI === Person.prototype.whoAmI) {
            throw new TypeError("You must implement the abstract method whoAmI");
        }
    }

    whoAmI() {
        // if this method is called from a subclass without being overwriten, throw a type error again
        throw new TypeError("You must implement the abstract method whoAmI");
    }
}

