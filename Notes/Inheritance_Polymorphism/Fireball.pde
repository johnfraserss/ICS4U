/** This class will be used for a Fireball
 * 
 * @since March 24, 2015
 * @version 1.0
 * @author Mr. Seidel
 */
class Fireball {
  PVector location;
  PVector velocity;
  PVector acceleration;
  
  /** Default constructor
   *
   */
  Fireball() {
    location = new PVector(random(width), random(height));
    velocity = new PVector(0, 0);
    acceleration = new PVector(0, 0);
  }
   
  /**
   * This function displays the Fireball
   */
   void display() {
     fill(#FA5353); //red
     ellipse(location.x, location.y, 50, 50);
   }
   
   /**
    * This function applies a force to the velocity vector
    *
    * @param force   force to be applied as a PVector
    */
   void applyForce(PVector force) {
     velocity.add(force);
     location.add(velocity);
   }
}
