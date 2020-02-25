// import files according to your own file structure
'use strict'
import Person from './Person';

class Employee extends Person {

    constructor(name, age) {
        this.name = name;
        this.age = age;
    }

    whoAmI() {
       return this.name + " " + this.age
    }
}