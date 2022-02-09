package test;

import java.util.*;

public class Sample {
	public static void parse(ArrayList<HashMap<String, String>> products) {
//		System.out.println(products);
		for(int i=0; i< products.size(); i++) {
			boolean flag = true;
			HashMap<String, String> product1 = products.get(i);
//			System.out.println(product1);
			if(product1.get("price").equals("600")) {
				System.out.println(product1.get("name"));
			}
			
		}
		
		for(HashMap<String, String> product1: products) {
			boolean flag = true;
//			HashMap<String, String> product1 = products.get(i);
//			System.out.println(product1);
			if(product1.get("price").equals("600")) {
				System.out.println(product1.get("name"));
			}
			
		}
	}
	
	public static void main(String[] args) {
		HashMap<String, String> product1 = new HashMap<String, String>();
		product1.put("name", "P1");
		product1.put("price", "600");
		
		HashMap<String, String> product2 = new HashMap<String, String>();
		product2.put("name", "P2");
		product2.put("price", "700");
		
		ArrayList<HashMap<String, String>> products = new ArrayList<HashMap<String, String>>();
		products.add(product1);
		products.add(product2);
		
		parse(products);
		
	}
	
	
}
