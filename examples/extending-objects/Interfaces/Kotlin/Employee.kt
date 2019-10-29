class Employee(firstName : String, lastName : String, age : Int, staffNumber : Int) : Person {
  override firstName : String 
  override lastName : String
  var age : Int
  var staffNumber : Int
    
  override fun toString() : String {
    return this.firstName + " " + this.lastName + " " + this.age + ", " + this.staffNumber;
  }
}