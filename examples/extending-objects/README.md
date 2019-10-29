# Extending Objects (Short Notes)

There are multiple ways to save time and force cohesion between different programmers: _Inheritance, Interfaces,_ and _Abstract Classes_.  Using these difference extensions, we can then use _Polymorphism_ to our advantage. Short description of each below.


| Type of Extension | Main Benefit | Language Specific Information |
| ----------------- | ------------ | ----------------------------- |
| **Interfaces** | Create a template of methods <br/><br/> Seen as a "contract" | C#, Java, Kotlin and Processing have support for this.  This is a work-around to the multiple-inheritance problem, and helps programmers to conform to a standard.  <br/><br/>Python and JavaScript do not support interfaces. |
| **Abstract Classes** | Like interfaces; however, you can create the body of functions and instance variables | C++, C#, Java, Kotlin, Python, and Processing all support abstract classes. <br/><br/> JavaScript does not support abstract classes (natively - it **can** be forced, but isn't a supported design pattern/data structure). |
| **Inheritance** | Full access to parent class | c++, C#, Java, Kotlin, Python, JavaScript and Processing all have support for this.  Out of the list, only Python allows for multiple-inheritance (usually more difficult to work with) |

# Polymorphism

Is the idea of using a single array/list to hold multiple different related data types.  See the examples in the Polymorphism folder above (not all languages are covered above).  

If you are using **JavaScript**, you can watch [this short video](https://www.youtube.com/watch?v=8a5BkwuZRK0) (watch from 0:00 to about 5:25) to get an idea of how it could work for you.  

If you are using **C++** you can find an [example here](http://www.cplusplus.com/doc/tutorial/polymorphism/).
