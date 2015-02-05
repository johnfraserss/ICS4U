/**
 * This Circle will be used as a generic outline for what to do when we want
 * do create a Circle on the screen.
 * 
 * @author Mr. Seidel
 * @version 1.0
 * @since Processing 2.1
 * @since Nov 2014
 */
class Circle {
    PVector location;
    float radius;
    
    /**
     * This function returns the diameter of this circle
     * 
     * @return the diameter of the circle
     */
    float diameter() {
        return radius * 2;
    }
    
    /**
     * This function returns the radius of this circle
     * 
     * @return the radius of the circle
     */
    float radius() {
        return radius;
    }
    
    /**
     * This function calculates and returns the circumference of this circle
     * 
     * @return the circumference of the circle (as a float)
     */
    float circumference() {
        return (float)(2 * Math.PI * radius);
    }
    
    /**
     * This function calculates and returns the area of this circle
     * 
     * @return the area of the circle (as a float)
     */
    float area() {
        return (float)(Math.PI * radius * radius);
    }
    
    /**
     * This function will calculate the distance between another 
     * point and this Circle
     * 
     * @param point the other point to calculate distance from (as PVector)
     * @return      the value of the distance between points (as a float)
     */
    float distanceFrom(PVector point) {
        return (dist(point.x, point.y, location.x, location.y));
    }
}
