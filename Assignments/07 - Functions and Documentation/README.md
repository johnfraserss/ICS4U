Task(s)
-------
Create a basic program that contains the following functions and uses them in a basic visual

```java
/*
If you are doing Processing, use floats
If you are doing P5.js, you just need to use vars and ensure you return a number (or array of numbers)

Below it is written as if you were doing the
assignment using Processing.

*/
addition(float, float) //returns a float
subtract(float, float) //returns a float
multiply(float, float) //returns a float
divide(float, float) //returns a float
maximum(float, float) //returns the larger of the two values as a float
minimum(float, float) //returns the smallest of the two values as a float
reverseArray(float []) //returns a reversed float [ ] array (can use ArrayLists if you prefer), program the algorithm yourself - do not use a built in "reverse()" function 
```

Below is an example of how the _addition_ function would be written:
```java
//Java (javadocs)

/** This function takes in two float values and returns the sum of those two values
 *
 *  @param first   This is the first float of the sum
 *  @param second  This is the second float of the sum
 *  @return        The sum of first and second as a float value
 */
float addition(float first, float second) {
	return first + second;
}


//JavaScript (JSdoc)
/** This function takes in two float values and returns the sum of those two values
 *
 *  @param {number} first   This is the first float of the sum
 *  @param {number} second  This is the second float of the sum
 *  @return {number}        The sum of first and second as a float value
 */
function addition(first, second) {
	return first + second;
}

```

Submission Item(s)
------------------
A basic working program that demonstrates your knowledge of the functions given above (including documentation) - do __not__ submit to the IN folder.  Create a repository online to hold your programming assignments.  Within that repository, you can create separate folders for each of the topics.  For example:

\| Home folder  
\|   
\| \- /Assignment 1  
\| \- /Assignment 2   
\| \- /etc.  

As long as your assignment is committed to the repository by the due date, it will be considered on time.

Include a _README.md_ file in the folder to explain what your program does.

[See this repository for an example](https://github.com/mrseidel-classes/sample-4U-code)



Due Date(s)
-------------
To be uploaded to the repository by October 4th, 2016

| | 0 | 1 | 2 | 3 | 4 |
|---| --- | --- | --- | --- | --- |
|A4.1: Work independently, using support documentation to resolve syntax issues during software development  | | | | | |  
|A4.3: Create fully documented (and well-named) program code according to industry standards (**javadocs/jsdocs**, block comments, line comments, variable names, etc.)  | | | | | |
