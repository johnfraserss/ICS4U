/**
 * This is a circle that bounces around the screen
 *
 * @class
 * @constructor
 * @param {number} r is a red value
 * @param {number} g is a green value
 * @param {number} b is a blue value
 */
function Circle(r, g, b) {
  this.x = width/2;
  this.y = 200;
  this.xspeed = random(-20, 20);
  this.yspeed = random(-20, 20);
  this.r = r;
  this.g = g;
  this.b = b;

  /**
   * Moves the circle around the screen
   */
  this.move = function() {
    this.x += this.xspeed;
    this.y += this.yspeed;
  }

  /**
   * Checks to see if the circle has hit 
   * the edge of the canvas
   */
  this.edges = function() {
    if (this.x < 0 || this.x > width) {
      this.xspeed *= -1;
    }
    if (this.y < 0 || this.y > height) {
      this.yspeed *= -1;
    }
  }

  /**
   * Draws the circle on the screen
   */
  this.render = function() {
    fill(this.r, this.g, this.b);
    ellipse(this.x, this.y, 10, 10);
  }
}
