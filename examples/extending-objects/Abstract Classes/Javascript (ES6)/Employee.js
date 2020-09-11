// import files according to your own file structure
import Person from './Person';

class Employee extends Person {

    constructor(name, age) {
        super()
        this.name = name;
        this.age = age;
    }

    whoAmI() {
       return this.name + " " + this.age
    }
}