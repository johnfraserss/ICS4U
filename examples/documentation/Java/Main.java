class Main {
  public static void main(String[] args) {
    System.out.println("Hello world!");
    System.out.println(inchesToCentimeters(1));
  }

  /**
  * Converts inches to centimeters
  * 
  * @param length - The length of measurement to convert
  *
  * @returns The converted value
  */
  private static float inchesToCentimeters(float length) {
    return length * 2.54f;
  }
}