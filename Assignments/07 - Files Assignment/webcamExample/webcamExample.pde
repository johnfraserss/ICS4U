import processing.video.*; //imports files, classes, etc. for using the webcam

Capture cam; //variable used to access webcam

void setup() {
  size(640, 480);

  String[] cameras = Capture.list(); //gets a list of video sources available

  if (cameras.length == 0) {
    println("There are no cameras available for capture.");
    exit();
  } 
  else {
    println("Available cameras:");
    for (int i = 0; i < cameras.length; i++) {
      println(cameras[i]);
    }

    // The camera can be initialized directly using an 
    // element from the array returned by list():
    cam = new Capture(this, cameras[0]);
    cam.start();
  }
}

void draw() {
  if (cam.available() == true) {
    cam.read();
    loadPixels(); // For OUTPUT window's pixels[] array

    for (int y=0; y < height; y++)
    {
      for (int x = 0; x < width; x++)
      {
        int i = x + y * width; //get current pixel

        // Pixel from cam (red, green, blue) or brightness
        float r = red(cam.pixels[i]);
        float g = green(cam.pixels[i]);
        float b = blue(cam.pixels[i]);
        float bright = brightness(cam.pixels[i]); //gets brightness of pixel
      
        //do stuff with the information from the pixel array

        // "Safety" code - ensures numbers are between 0 and 255
        r = constrain(r, 0, 255);
        g = constrain(g, 0, 255);
        b = constrain(b, 0, 255);

        pixels[i] = color(r, g, b);
      }
    }

   updatePixels(); // OUTPUT window's pixels[] array
  }
}

