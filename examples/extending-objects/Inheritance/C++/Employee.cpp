#include <string>
using namespace std;
#include "Person.cpp"

class Employee : public Person {
  private:
	 int staffNumber;
  
  public:
	 Employee(string firstName, string lastName, int age, int staffNumber) : Person(firstName, lastName, age) {
		this->staffNumber = staffNumber;
	}
  
	 string toString() {
		return Person::toString() + ", " + to_string(this->staffNumber);
	}
};