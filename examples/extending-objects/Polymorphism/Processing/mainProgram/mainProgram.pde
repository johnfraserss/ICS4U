Person [] people = new Person[2];

void settings() {}
void setup() {
  people[0] = new Person("Marge", "Simpson", 36);
  people[1] = new Employee("Homer", "Simpson", 28, 1007);

  for (Person p : people) {
    print(p);
  }
}
void draw() {
  noLoop();
}
