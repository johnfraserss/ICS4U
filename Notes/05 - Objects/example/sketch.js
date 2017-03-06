var circles = [];
var circle;

function setup() {
  createCanvas(500, 500);
  for (var i = 0; i < 100; i++) {
    circle = new Circle(random(255), random(255), random(255));
    circles.push(circle);
  }

  /*  //This is a basic look at
      //the other "JSON"-style object
  circle = {
    x : 50,
    y : 100,

    move : function() {
       this.x += 1;
       this.y += 1;
    }
  }
  */
}

function draw() {
  background(0);
  for (var i = 0; i < circles.length; i++) {
    circles[i].render();
    circles[i].move();
    circles[i].edges();
  }
}
