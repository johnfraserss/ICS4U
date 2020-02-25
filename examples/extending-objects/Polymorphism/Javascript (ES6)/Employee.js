// assuming import of the Person class (Person.js) - Read up on exports/imports in ES6 if you are confused

class Employee extends Person {
    constructor(age, name, employeeNum, tenure) {
        super(age, name);
        this.employeeNum = employeeNum;
        this.tenure = tenure;
    }
    
    toString() {
        return (this.age + ", " + this.name + ", " + this.employeeNum + ", " + this.tenure)
    }
}