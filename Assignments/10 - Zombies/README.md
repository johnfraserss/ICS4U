Task(s)
-------
**NOTE: Due to time constraints, this assignment has been cancelled**

You **may** work in groups of 2 to complete this assignment

You will need PVectors, Objects, Inheritance, Files, Sorting, and Searching knowledge to complete this assignment.

This is a simulated game, there is no input from the user after the program starts, as it should be completely automated.

* Run the program OUT\SEIDEL\ICS4U\ZOMBIES\RunMe.exe to generate a .txt file for yourself - everyone's will be different, but your program will have to work with any of the .txt files generated!
* This text file will hold information regarding:
  * The minimum size of your screen (Width: 600-1000, Height: 400-700)
  * Zombie spawn interval (in seconds) (10-40 seconds)
  * Zombie locations (4-8 locations)
  * Maximum amount of Zombies allowed on the screen (1-8)
  * Fireman spawn interval (in seconds) (60-90 seconds)
  * Maximum amount of Firemen allowed on the screen (2-4)
  * Number of Fireman success tokens (3-8)
  * Fire Station locations (1-4)

* Properties of Zombies:
  * Zombies are 10x10 rectangular beings that leave a trail of fire (4x4 rectangle) behind them that is no longer than 10 fire-trail-lengths.
  * Zombies cannot come within 50 pixels of the centre of a fire station
  * They move randomly via a noise() function, unless within 20 pixels of a Fireman.  
  * If they leave the screen completely a new Zombie gets created at one of the spawn locations in the text file.
  * Zombies move 20% faster than Firemen

* Properties of Fireman
  * Firemen spawn at fire stations (16x16 yellow block)
  * Firemen dowse fire trails as they pass over them
  * Firemen are blue 6x6 rectangular beings that go towards the __closest__ Zombie that is **not** in combat and attempts to drown them (moves towards them if they are within a range of 50 pixels, otherwise they randomly roam)!
  * If successful, they then take the time to move around cleaning up the fire trail that the zombie left behind - always working on the __closest__ piece of fire trail.  
  * If unsuccessful, the Fireman turns into a Zombie (this **can** override the maximum Zombie count temporarily)!

* Other items of note:
  * The fight between a Fireman and a Zombie takes at least 2 seconds to complete.
  * Only 1 Fireman can fight a Zombie.  If a Zombie is in combat, then it is not considered the __closest__ Zombie.
  * The success rate of a Fireman is determined by the number of _success tokens_ they have.  Each token is worth 5% to a maximum of 85% success rate. (i.e. if they have 10 _success tokens_ they have a 50% chance of succeeding in the fight against a Zombie)
  * For every successful fight, that particular Fireman gains another _success token_

* From a programming stand point:
  * Firemen inherit properties and methods from Zombies
  * You should always have a sorted arrays of Zombies and Firemen (you need to figure out how to do this so it is effective)
  * The __closest__ Zombie is always found via a binary search on positions (you need to figure out how to do this so it is effective)

Submission Item(s)
------------------
Hand in the work to the IN folder

Due Date(s)
-----------
This assignment is due by May 4th by 3:30pm - **no exceptions**

__Note__: The long-term assignment is also due this day

Rubric(s)
---------
This rubric is out of a total of 28

| | 0 | 1 | 2 | 3 | 4 |
|---| --- | --- | --- | --- | --- |
|A1.2: Demonstrate type conversion (string-to-integer, string-to-float, etc.)  | | | | | |
|A2.2: Use modular design concepts that support reusable code (e.g. encapsulation, inheritance, polymorphism, etc.)  | | | | | |
|A3.2: Create and use a searching algorithm on an array of of sorted objects  | | | | | |
|A3.4: Create and use a sorting algorithm on an array of objects  | | | | | |
|A4.3: Create fully documented program code according to industry standards (e.g. javadocs, block comments, line comments, variable naming practices, etc.)  | | | | | |
|C1.2: Demonstrate the ability to apply data encapsulation in program design  | | | | | |
|C1.3: Demonstrate the ability to apply the process of functional decomposition in subprogram design  | | | | | |

The final mark for this assignment will be calculated as: __Mark = R * C__ where **R = Rubric total** and **C = Completed the task as requested/described (worth 0 or 1)**
