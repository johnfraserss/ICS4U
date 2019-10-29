// assuming include of <string> and using the same namespace

class Person {
  protected:
	string firstName;
	string lastName;
	int age;
  
  public:
	Person(string firstName, string lastName, int age) {
		this.firstName = firstName;
		this.lastName = lastName;
		this.age = age;
	}
	
	~Person();
	
    string toString() {
		return this.firstName + " " + this.lastName + ", " + this.age;
	}
};