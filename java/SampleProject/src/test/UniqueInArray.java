package test;

import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;

public class UniqueInArray {
	public static void main(String[] args) {
		int[] a = new int[6];
		a[0] = 2;
		a[1] = 2;
		a[2] = 1;
		a[3] = 2;
		a[4] = 3;
		a[5] = 1;
		HashMap<Integer, Integer> hm = new HashMap<Integer, Integer>();
		for(int i=0; i< a.length; i++) {
			hm.put(a[i], hm.getOrDefault(a[i], 0) + 1);
		}
		
		Iterator<Map.Entry<Integer, Integer>> it = hm.entrySet().iterator();
	    while (it.hasNext()) {
	        Map.Entry<Integer, Integer> pair = (Map.Entry<Integer, Integer>)it.next();
	        System.out.println(pair.getKey() + " = " + pair.getValue());
	        if((int)pair.getValue() == 1) {
	        	System.out.println(pair.getKey());
	        }
	    }
	}
}
