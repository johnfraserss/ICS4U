var capture;
// var pixR = [];
// var pixG = [];
// var pixB = [];
// var pixA = [];

function setup() {
  createCanvas(640, 480);
  capture = createCapture(VIDEO);
  capture.size(64, 48);
  //capture.hide();   //hides the video feed
  //pixelDensity(1);  //used for retina displays to compensate
  noStroke();
}

function draw() {
  background(255);
  capture.size(64, 48);
  capture.loadPixels();
  for (var x = 0; x < capture.width; x++) {
    for (var y = 0; y < capture.height; y++) {
      // pixR.push(capture.get(x, y)[0]);
      // pixG.push(capture.get(x, y)[1]);
      // pixB.push(capture.get(x, y)[2]);
      // pixA.push(capture.get(x, y)[3]);

      //sometimes the "pixels" are images - in that case
      //we don't want to do anything with it until the
      //video has settled into pixels.
      if (!(capture.get(x,y) instanceof p5.Image)) {
        fill(capture.get(x,y)[0],
             capture.get(x,y)[1],
             capture.get(x,y)[2],
             capture.get(x,y)[3]);
        ellipse(x * 10, y * 10, 8, 8);
      }
    }
  }
}
