abstract class Person {
  String firstName;
  String lastName;
  int age;
  
  Person(String firstName, String lastName, int age) {
    this.firstName = firstName;
    this.lastName = lastName;
    this.age = age;
  }
  
  String toString() {
    return this.firstName + " " + this.lastName + ", " + this.age;
  }
  
  abstract void increaseAge();
}