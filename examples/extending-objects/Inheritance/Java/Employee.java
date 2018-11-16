public class Employee extends Person {
  private int staffNumber;
  
  public Employee(String firstName, String lastName, int age, int staffNumber) {
    super(firstName, lastName, age);
    this.staffNumber = staffNumber;
  }
  
  public String toString() {
    return super.toString() + ", " + this.staffNumber;
  }
}