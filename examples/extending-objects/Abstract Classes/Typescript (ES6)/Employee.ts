// import files according to your own file structure
import Person from './Person';

class Employee extends Person {
    name;
    age;

    constructor(name: string, age: number) {
        super()
        this.name = name;
        this.age = age;
    }

    whoAmI() {
       return this.name + " " + this.age
    }
}

export default Employee