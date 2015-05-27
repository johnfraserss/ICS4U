/** This class will be used for a Waterball
 * 
 * @since March 24, 2015
 * @version 1.0
 * @author Mr. Seidel
 */
class Waterball extends Fireball {
  Waterball() {
    super();
  }
  void display() {
    fill(#7787F0); //blue
    rect(location.x, location.y, 50, 50);
  }
}

