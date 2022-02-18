'''
Recursive example using Turtle graphics to draw

Modified from this work
https://p5js.org/examples/simulate-recursive-tree.html
'''


def branch(height, theta): # recursive function
    '''
    Draws a single branch of a tree.

    Parameters
    ----------
    height : int
      The length of the branch in the tree
    theta : int
      The angle between the branches

    '''
    height *= 0.66;

    if (height > 10): # if this is true, continue recursion
        right(theta)
        forward(height)
        branch(height, theta) # recursive call with new height, same theta

        backward(height)
        left(theta)

        left(theta)
        forward(height)
        branch(height, theta) # recursive call with new height, same theta

        backward(height)
        right(theta)

    return
  

from turtle import *
speed(10)
color('black')
left(90)
forward(120) # draws initial trunk
branch(120, 45) # change second parameter for angle
backward(120)
done()