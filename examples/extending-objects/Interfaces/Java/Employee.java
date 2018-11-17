public class Employee implements Person {
  private String firstName, lastName;
  private int age, staffNumber;
  
  public Employee(String firstName, String lastName, int age, int staffNumber) {
    this.firstName = firstName;
    this.lastName = lastName;
    this.age = age;
    this.staffNumber = staffNumber;
  }
  
  public String toString() {
    return this.firstName + " " + this.lastName + ", " + age + ", " + this.staffNumber;
  }
}