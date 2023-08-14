public class variables {
	public static void main(String[] args) {
		System.out.println(1);
		System.out.println(3);
		System.out.println(2*3);
		/* BODMAS rule. 
		Will be multiplied and printed */
		System.out.println("Multiplication" + 4 * 3);
		// Will not add and print as it is
		System.out.println("Addition" + 2 + 4);
		// Will work as expected. Addition and print
		System.out.println("Multiplication" + (3 +4));

		String name = "John";
		System.out.println(name);

		int myNum = 15;
		System.out.println(myNum);

		int myNum1;
		myNum1 = 21;
		System.out.println(myNum1);
		myNum1 = 25;
		System.out.println(myNum1);

		final int myNum2;
		myNum2 = 20;
		System.out.println(myNum2);
		// myNum2 = 45;
		// System.out.println(myNum2);
		String carName = "volvo";
		System.out.println(carName);

		String firstname = "Kartik";
		String lastname = "Iyyer";
		String fullname = firstname + " " + lastname;
		String middlename = new String("S");
		System.out.println(fullname);
		System.out.println(middlename);

		int x = 1;
		int y = 2;
		int sum = x + y;
		System.out.println(x+y);
		System.out.println(sum);
		int x1=6,y1=5,z1=5;
		int a,b,c;
		a=b=c=10;
		System.out.println(x1 + y1 + z1);
		System.out.println(a+b+c);

		byte small = 10;
		System.out.println(small);

		double big = 10000.48484848;
		System.out.println(big);

		long medium = 1003L;
		System.out.println(medium);

		// int sum1 = medium + small;
		// System.out.println(sum1);

		// need to specify f or d at the end if e is used.
		float expo = 2e4f;
		System.out.println(expo);

	}
}