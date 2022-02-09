import java.util.*;
import java.util.Arrays;

public class LongestPrefix {
	static String findLongest(String str1, String str2) {
		String substr = new String();
		for(int i=0,j=0; i<str1.length() && j<str2.length(); i++,j++) {
			if(str1.charAt(i) != str2.charAt(j)) {
				break;
			}
			substr += str1.charAt(i);
		}
		return substr;
	}
	public static void main(String[] args) {
		String arr[] = {"geek", "geeks", "gee", "geedfd"};
		Arrays.sort(arr);
		for (String a : arr) {
			System.out.println(a);
		}
		HashMap<String,String> map = new HashMap<String,String>();
		map.put("asdf", "sadf");
		System.out.println(map.get("asdf"));
		for(Map.Entry<String, String> entry: map.entrySet()) {
			System.out.println(entry.getKey());
		}
		Map<String,String> map1 = Collections.emptyMap();
		
		System.out.println(map1);
		System.out.println(findLongest(arr[0], arr[arr.length-1]));
	}
}