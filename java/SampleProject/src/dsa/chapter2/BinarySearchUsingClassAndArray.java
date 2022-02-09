package dsa.chapter2;


import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;

class Person {
	String firstName;
	String lastName;
	int age;
	
	public Person(String firstName, String lastName, int age) {
		this.firstName = firstName;
		this.lastName = lastName;
		this.age = age;
	}
	
	public String getLastName() {
		return lastName;
	}
}

class OrderedAtt {
	Person[] persons;
	int numberOfStudentsPresent;
	public OrderedAtt(int numberOfStudents) {
		persons = new Person[numberOfStudents];
		numberOfStudents = 0;
	}
	
	public void addStudent(Person person) {
		
	}
}


public class BinarySearchUsingClassAndArray {
	public static void main(String[] args) {
		Logger logger = LogManager.getLogger();
//		System.out.println("Binary search using Class and Array.");
		logger.info("Binary search using Class and Array.");
	}
}
