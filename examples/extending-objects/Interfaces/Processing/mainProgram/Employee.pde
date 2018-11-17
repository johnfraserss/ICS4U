class Employee implements Person {
  String firstName, lastName;
  int age, staffNumber;
  
  Employee(String firstName, String lastName, int age, int staffNumber) {
    this.firstName = firstName;
    this.lastName = lastName;
    this.age = age;
    this.staffNumber = staffNumber;
  }
  
  String toString() {
    return this.firstName + " " + this.lastName + ", " + age + ", " + this.staffNumber;
  }
}