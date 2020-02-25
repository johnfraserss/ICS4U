class Person {
    constructor(first, last, age) {
        self.firstname = first
        self.lastname = last
        self.age = age
	}

    whoAmI() {
        return self.firstname + " " + self.lastname + ", " + self.age.toString()
	}
}