#include <string>
using namespace std;

class Person {
  protected:
  	string firstName;
  	string lastName;
  	int age;
    
  public:
  	Person(string firstName, string lastName, int age) {
  		this->firstName = firstName;
  		this->lastName = lastName;
  		this->age = age;
  	}
	
    string toString() {
		return this->firstName + " " + this->lastName + ", " + to_string(this->age);
	}
};