public class StringReverse {
	static String recursion(String str) {
		if(str.isEmpty()) {
			return "";
		}
		return recursion(str.substring(1)) + str.charAt(0);
	}

	static void normal(String str) {
		String reversedStr = "";
		for(int i=str.length()-1;i>=0;i--) {
			reversedStr += str.charAt(i);
		}
		System.out.println(reversedStr);
	}
	public static void main(String[] args) {
		String str = "This is sample file";
		System.out.println(recursion(str));

		normal(str);
	}
}