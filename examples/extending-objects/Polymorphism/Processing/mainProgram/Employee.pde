class Employee extends Person {
  int staffNumber;
  
  Employee(String firstName, String lastName, int age, int staffNumber) {
    super(firstName, lastName, age);
    this.staffNumber = staffNumber;
  }
  
  String toString() {
    return super.toString() + ", " + this.staffNumber;
  }
}