# Example retrieved from https://www.python-course.eu/python3_abstract_classes.php
# Date retrieved: Nov 15, 2018
# Example slightly modified


from abc import ABC, abstractmethod

# as an abstract class, you cannot instantiate this class anymore
class Person(ABC):
    def __init__(self, first, last, age):
        self.firstname = first
        self.lastname = last
        self.age = age
        super().__init__()

    def __str__(self):
        return self.firstname + " " + self.lastname + ", " + str(self.age)
		
    @abstractmethod
    def incrementAge(self):
        return
