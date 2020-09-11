/// import the Person and Employee classes accordingly

// this will throw a type error as Person is an abstract class and cannot be directly instantiated
const marge = new Person();

// This will work as Employee extends (derives) the abstract class Person and implements its methods
const homer = new Employee("Homer", "28");
homer.whoAmI();
