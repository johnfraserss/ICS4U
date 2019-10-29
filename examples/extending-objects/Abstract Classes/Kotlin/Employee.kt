// Assuming everything is in the same package

class Employee(firstName : String, lastName : String, age : Int, staffNumber : Int) : Person(firstName, lastName, age) {
  
  override fun toString() : String {
    return super.toString() + ", " + this.staffNumber;
  }
}