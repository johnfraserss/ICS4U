/**
 * This Circle will be used as a generic outline for what to do when we want
 * to create a circle on the screen.
 * 
 * @author Mr. Seidel
 * @version 1.2
 * @since Processing 2.1
 * @since Nov 2014
 */
class Circle {
  PVector location; 
  float radius;

  /**
   * This is a constructor for the Circle
   * 
   * @param location  this is a passed in initial location for the Circle object
   * @param radius this is a passed in initial radius for the Circle object
   */
  Circle(PVector location, float radius) {
    this.location = location;
    this.radius = radius;
  }

  /**
   * This is the default constructor.  It uses the Circle(PVector, float) constructor to
   * generate a default circle at (0, 0) with a radius of 0
   *
   */
  Circle() {
    this(new PVector(0, 0), 0);
  }

  /**
   * This function returns the diameter of this circle
   * 
   * @return the diameter of this Circle object (as a float)
   */
  float diameter() {
    return radius * 2;
  }

  /**
   * This function returns the radius of this circle
   * 
   * @return the radius of this Circle object (as a float)
   */
  float radius() {
    return radius;
  }

  /**
   * This function calculates and returns the circumference of this circle
   * --This function uses the diameter() function 
   *
   * @return the circumference of this Circle object (as a float)
   */
  float circumference() {
    return (float)(Math.PI * diameter());
  }

  /**
   * This function calculates and returns the area of this Circle object
   * 
   * @return the area of the circle (as a float)
   */
  float area() {
    return (float)(Math.PI * radius * radius);
  }

  /**
   * This function will calculate the distance between another 
   * point and this Circle object
   * 
   * @param point the other point to calculate distance from (as PVector)
   * @return      the value of the distance between points (as a float)
   */
  float distanceFrom(PVector point) {
    return (float)(dist(point.x, point.y, location.x, location.y));
  }

  /**
   * This function will display the Circle object on the screen
   * --This function uses the local diameter() function
   */
  void display() {
    ellipse(location.x, location.y, diameter(), diameter());
  }
}
