PImage img;

void setup() {
  size(400, 400);
  img = loadImage("D:\\School\\smiley.png");
}

void draw() {
  background(0);
  image(img, 0, 0);
}
