package test;

import com.google.gson.Gson;

class Person {
	String name;
	String sport;
	int age;
	int id;
	int[] lastScores;
	public Person() {
		
	}
//	public Person(String name, String sport, int age, int id, int[] lastScores) {
//		this.name = name;
//		this.sport = sport;
//		this.age = age;
//		this.id = id;
//		this.lastScores = lastScores;
//	}
}

public class JSONParser {
	public static void main(String[] args) {
		String personJSON = "{" + 
				"  \"name\" : \"Ronaldo\"," + 
				"  \"sport\" : \"soccer\"," + 
				"  \"age\" : 25," + 
				"  \"id\" : 121," + 
				"  \"lastScores\" : [ 2, 1, 3, 5, 0, 0, 1, 1 ]" + 
				"}";
		Gson g = new Gson();
		Person person = g.fromJson(personJSON, Person.class);
		
		System.out.println(person.age);
	}
}
