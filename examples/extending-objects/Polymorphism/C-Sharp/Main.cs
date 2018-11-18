class Main {
	
	Person [] employees = new Person[2];
	
	public static void Main(string[] args) {
		employees[0] = new Person("Marge", "Simpson", 36);
		employees[1] = new Employee("Homer", "Simpson", 28, 1007);
	
		foreach (Person p in employees) {
			Console.WriteLine(p);
		}
	}
}