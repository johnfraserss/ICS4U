public class Employee : Person {
  private firstName, lastName;
  private int age, staffNumber;
  
  Employee(String firstName, String lastName, int age, int staffNumber) {
    this.firstName = firstName;
    this.lastName = lastName;
    this.age = age;
    this.staffNumber = staffNumber;
  }
  
  public String ToString() {
    return this.firstName + " " + this.lastName + ", " + age + ", " + this.staffNumber;
  }
}