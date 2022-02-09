package test;

class AdditionOfNumbers {
	int sum;
	int a;
	public static int retSum(int a, int b)
	{
		AdditionOfNumbers ann = new AdditionOfNumbers();
		return ann.sum(a, b);
	}
	
	public int sum(int a, int b) {
		return a + b;
	}
	public static void main(String[] args)
	{
		int a=5,b=4,c;
		AdditionOfNumbers an = new AdditionOfNumbers();
		c = an.sum(a, b);
		System.out.println(c);
		
		
	}
}
