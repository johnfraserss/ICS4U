public class Main {
	Person [] employees = new Person[2];
	
	public static void main(String args[]) {
		employees[0] = new Person("Marge", "Simpson", 36);
		employees[1] = new Employee("Homer", "Simpson", 28, 1007);
	
		for (Person p : employees) {
			System.out.println(p);
		}
	}
}