# Example retrieved from https://www.python-course.eu/python3_inheritance.php
# Date retrieved: Nov 15, 2018
# Example slightly modified

class Person:
    def __init__(self, first, last, age):
        self.firstname = first
        self.lastname = last
        self.age = age

    def __str__(self):
        return self.firstname + " " + self.lastname + ", " + str(self.age)
