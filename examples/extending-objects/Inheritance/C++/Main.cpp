#include <iostream>
using namespace std;
#include "Employee.cpp"

int main(void) {
	Person marge("Marge", "Simpson", 36);
	Employee homer("Homer", "Simpson", 28, 1007);
	
	cout << marge.toString() << "\n" << endl;
	cout << homer.toString() << "\n" << endl;
}