# Example retrieved from https://www.python-course.eu/python3_inheritance.php
# Date retrieved: Nov 15, 2018
# Example slightly modified

# More examples available here: https://pythonspot.com/polymorphism/

from person import Person
from employee import Employee

employees = []

employees[0] = Person("Marge", "Simpson", 36)
employees[1] = Employee("Homer", "Simpson", 28, "1007")

for item in employees:
    print(item)
