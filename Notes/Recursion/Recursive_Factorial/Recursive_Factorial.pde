void setup() {
  size(600, 600);
}


void draw() {
  //get a value to use for factorial functions
  int n = int(map(mouseX, 0, width, 1, int(width/100) + 1));

  //display colour gradient
  noStroke();
  for (int i = 0; i < int(width/100) + 1; i++) {
    fill(255/(i + 1));
    rect(i * 100, 0, (i + 1)*100, height);
  }
  
  //display text
  textSize(50);
  fill(255);
  text(factorial(n), width/2, height/2);
  //text(recursiveFactorial(n), width/2, height/2);
}


int factorial(int n) {
  int answer = 1;
  for (int i = 1; i <= n; i++) {
    answer *= i;
  }
  return answer;
}

int recursiveFactorial(int n) {
  if (n <= 1) {
    return 1;
  }
  return n * recursiveFactorial(n - 1);
}
