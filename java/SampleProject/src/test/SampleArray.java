package test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.Map;

public class SampleArray {
	public static void main(String[] args) {
		ArrayList<String> users = new ArrayList<String>();
		users.add("John");
		
		// Create this cc to users
		HashMap<Integer, Integer> hm = new HashMap<Integer, Integer>();
		hm.put(0, 0);
		hm.put(1, 1);
		hm.put(1, 2);
		
		for (Map.Entry<Integer, Integer> entry : hm.entrySet()) {
		    Integer key = entry.getKey();
		    Integer value = entry.getValue();
		    System.out.println(key + " " + value);
		}
		
		
	}
}
