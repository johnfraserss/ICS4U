// Assuming everything is in the same package

abstract class Person (firstName : String, lastName : String, age : Int) {
  val firstName : String
  val lastName : String
  var age : Int

  init {
	this.firstName = firstName
	this.lastName = lastName
	this.age = age   
  }
  
  abstract fun toString() : String {
    return this.firstName + " " + this.lastName + ", " + this.age;
  }
}