#include <iostream>

/**
  * Converts inches to centimeters
  * 
  * @param length - The length of measurement to convert
  *
  * @returns The converted value
  */
float inchesToCentimeters(float length) {
    return length * 2.54f;
}

int main() {
  std::cout << "Hello World!\n";
  std::cout << inchesToCentimeters(1);
}


