import ddf.minim.*;

Minim minim;
AudioPlayer song;

void setup() {
  size(800, 250);
  minim = new Minim(this);
  song = minim.loadFile("D:\\School\\hourLong.mp3");
  song.play();
}

void draw() {
  background(0);
  textSize(40);
  text("Left Channel Value: " + song.left.get(0), 10, 50);
  text("Right Channel Value: " + song.right.get(0), 10, 100);
}

