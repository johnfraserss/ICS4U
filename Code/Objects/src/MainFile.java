/**
 * 
 */
package objectspkg;

/**
 * This file is to test the running of the work you completed in your shapes.
 * 
 * @author Mr. Seidel
 * @since Sept 16, 2015
 * @since JDK 8
 * @version 1.0
 */
public class MainFile {

	/**
	 * The main runnable area of the whole program 
	 * (like the "setup()" and "draw()" functions in Processing)
	 * 
	 * @param args Command-line arguments (not used in this case)
	 */
	public static void main(String[] args) {
		
		double radius = 5.0;
		double width = 5;
		double height = 10;
		
		SVector vertex = new SVector(50, 50);
		
		Circle c = new Circle();
		Circle c2 = new Circle(radius);
		Circle c3 = new Circle(radius, vertex));
		
		Rectangle r = new Rectangle();
		Rectangle r2 = new Rectangle(vertex);
		Rectangle r3 = new Rectangle(vertex, width, height);
		
		System.out.println("Circle's radius = "   + c.radius());
		System.out.println("Circle's position = " + c.centre());
		System.out.println("Circle's perimeter = "   + c.perimeter());
		System.out.println("Circle's area = "   + c.area());
		
		System.out.println("Circle2's radius = "   + c2.radius());
		System.out.println("Circle2's position = " + c2.centre());
		System.out.println("Circle2's perimeter = "   + c2.perimeter());
		System.out.println("Circle2's area = "   + c2.area());
		
		System.out.println("Circle3's radius = "   + c3.radius());
		System.out.println("Circle3's position = " + c3.centre());
		System.out.println("Circle3's perimeter = "   + c3.perimeter());
		System.out.println("Circle3's area = "   + c3.area());
		
		
		
		/*
		 * Complete the remainder of the println statements in order to show that your other
		 * work has been completed properly
		 */
		
	}
}
		
		
		
		
		
		
		
		
		
		