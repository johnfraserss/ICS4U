//Assuming the import of the person.js file is available

class Employee extends Person {

    constructor(first, last, age, staffnum) {
        super(first, last, age)
        self.staffnumber = staffnum
	}
		
    whoAmI() {
        return this.first + " " + this.last + ", " +  self.staffnumber
	}
}