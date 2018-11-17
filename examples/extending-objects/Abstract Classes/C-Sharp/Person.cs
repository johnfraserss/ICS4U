public abstract class Person {
  protected String firstName { get; set; }
  protected String lastName { get; set; }
  protected int age { get; set; }
  
  public Person(String firstName, String lastName, int age) {
    this.firstName = firstName;
    this.lastName = lastName;
    this.age = age;
  }
  
  public String ToString() {
    return this.firstName + " " + this.lastName + ", " + this.age;
  }
  
  abstract public void IncreaseAge();
}