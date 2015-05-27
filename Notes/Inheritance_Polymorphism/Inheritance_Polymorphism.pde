Fireball f;
Waterball w; //inheritance
Fireball [] ammo = new Fireball[3]; //polymorph

void setup() {
  size(600, 600);
  w = new Waterball();
  f = new Fireball();
  ammo[0] = new Fireball();
  ammo[1] = new Fireball();
  ammo[2] = new Waterball();
}

void draw() {
  //f.display();
  //w.display();
  for (int i = 0; i < ammo.length; i++) {
    ammo[i].display();
  }
}

void mousePressed() {
 f.applyForce(new PVector(0, -1));
 w.applyForce(new PVector(-1, 0)); 
}
