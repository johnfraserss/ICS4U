Task(s)
-------
You **may** work in pairs for this assignment.

1. Read in information from the webcam
2. Create a selectionSort() function which uses the selection sort algorithm (create this yourself, do not use built-in functions to do all the heavy lifting).
3. On LEFT-CLICK, use your selectionSort() function to sort the pixels in __ascending__ order by brightness and display them accordingly
4. On RIGHT-CLICK, use your selectionSort() function to sort the pixels in __descending order__ by their red, green, **or** blue value (your choice).
5. In a block comment in your main file, compare the run time and worst-case computational complexity of the selectionSort() algorithm vs a merge sort algorithm.  See Notes section below for more info

Submission Item(s)
------------------
Hand in the work to the IN folder

Due Date(s)
-----------
This assignment is due by April 21st, 2015 at 3pm - **no exceptions**

Rubric(s)
---------
This rubric is out of a total of 12

| | 0 | 1 | 2 | 3 | 4 |
|---| --- | --- | --- | --- | --- |
|C2.3: Compare the efficiency of sorting algorithms using run times and computational complexity analysis  | | | | | |
|A4.1: Work independently, using support documentation to resolve syntax issues during software development  | | | | | |
|A4.3: Create fully documented (and well-named) program code according to industry standards (**javadocs**, block comments, line comments, variable names, etc.)  | | | | | |

The final mark for this assignment will be calculated as: __Mark = R * C__ where **R = Rubric total** and **C = Completed the task as requested/described (worth 0 or 1)**

Notes:
------
To roughly calculate the computational complexity of a sorting algorithm (also known as "Big-Oh" notation), you count the number of statements executed, the number of iterations of a loop, and/or the number of comparisons performed.

For example, the [Bubble Sort] (http://mathbits.com/MathBits/Java/arrays/Bubble.htm) algorithm runs in O(n^2) time on average, and as a worst-case (read as "Oh n squared"), because of the two nested _for_ loops, which runs through and compares every single item with all items previous to it.
