class Fibonacci {


	static int fibrec(int n) {
		if(n <=1) {
			return n;
		}
		return fibrec(n-1) + fibrec(n-2);
	}

	public static long multiplyNumbers(int num)
    {
        if (num >= 1)
            return num * multiplyNumbers(num - 1);
        else
            return 1;
    }

	public static void main(String[] args) {
		// 0, 1, 1, 2, 3, 5, 8

		int len = 6;
		int sum = 0;
		int i = 0;
		int j = 1;

		while(len >=2) {
			sum = i + j;
			// System.out.println(sum);
			i = j;
			j = sum;
			len--;
		}
		// System.out.println(sum);
		System.out.println(""+fibrec(6));


	}
}