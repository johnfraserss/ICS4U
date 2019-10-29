// assuming include of <string> and using the same namespace

class Employee : public Person {
  private:
	int staffNumber;
  
  public:
	Employee(string firstName, string lastName, int age, int staffNumber) {
		Person::(firstName, lastName, age);
		this.staffNumber = staffNumber;
	}
	
	~Employee();
  
	string toString() {
		return Person::toString() + ", " + this.staffNumber;
	}
};