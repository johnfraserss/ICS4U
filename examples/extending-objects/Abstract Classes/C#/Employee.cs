public class Employee : Person {
  private int staffNumber;
  
  Employee(String firstName, String lastName, int age, int staffNumber) {
    base(firstName, lastName, age);
    this.staffNumber = staffNumber;
  }
  
  public String ToString() {
    return base.ToString() + ", " + this.staffNumber;
  }
}