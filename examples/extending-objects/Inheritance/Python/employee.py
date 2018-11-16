# Example retrieved from https://www.python-course.eu/python3_inheritance.php
# Date retrieved: Nov 15, 2018
# Example slightly modified

from person import Person
class Employee(Person):

    def __init__(self, first, last, age, staffnum):
        super().__init__(first, last, age)
        self.staffnumber = staffnum

    def __str__(self):
        return super().__str__() + ", " +  self.staffnumber

