package dsa.chapter2;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

class OrderedAttendances {
	int[] studentsInAscendingOrder;
	int numberOfStudentsPresent;
	
	public OrderedAttendances(int numberOfStudents) {
		studentsInAscendingOrder = new int[numberOfStudents];
		numberOfStudentsPresent = 0;
	}
	
	public void addStudent(int rollNumber) {
		long startTime = getCurrentTimeInMillis();
		System.out.println("Start time: " + startTime);
		System.out.println("Adding student with roll number: " + rollNumber);
		if(numberOfStudentsPresent==0) {
			studentsInAscendingOrder[numberOfStudentsPresent] = rollNumber;
		} else {
			int i;
			for(i=0; i<numberOfStudentsPresent; i++) {
				if(rollNumber < studentsInAscendingOrder[i]) {
					break;
				}
			}
			for(int j=numberOfStudentsPresent; j>i; j--) {
				studentsInAscendingOrder[j] = studentsInAscendingOrder[j-1];
			}
			
			studentsInAscendingOrder[i] = rollNumber;
		}
		numberOfStudentsPresent++;
		
		long endTime = getCurrentTimeInMillis();
		System.out.println("End time: " + endTime);
		System.out.println("Time taken: " + (endTime-startTime));
	}
	
	private long getCurrentTimeInMillis() {
		return System.currentTimeMillis();
	}
	
	public void searchStudent(int rollNumber) {
		int minLimit = 0;
		int maxLimit = numberOfStudentsPresent;
		int midLimit;
		int numberOfSearches = 1;
		while(minLimit <= maxLimit) {
			midLimit = (minLimit+maxLimit) / 2;
			System.out.println("Current middle: " + midLimit);
			if(rollNumber == studentsInAscendingOrder[midLimit]) {
		System.out.println("Student found with roll number: " + rollNumber + " in " + numberOfSearches + " searches.");
				break;
			} else if(rollNumber > studentsInAscendingOrder[midLimit]) {
				minLimit = midLimit+1;
			} else {
				maxLimit = midLimit-1;
			}
			numberOfSearches++;
		}
		
		if(minLimit > maxLimit) {
			System.out.println("Student with roll number: " + rollNumber + " not found.");
		}
	}
	
	public void displayCurrentAttendance() {
		long startTime = getCurrentTimeInMillis();
		System.out.println("Students present in the class in ascending order.");
		for(int i=0; i<numberOfStudentsPresent; i++) {
			System.out.println(studentsInAscendingOrder[i]);
		}
		long endTime = getCurrentTimeInMillis();
		System.out.println("Time taken: " + (endTime-startTime));
	}
}

public class BinarySearchUsingArray {
	public static void main(String args[]) throws NumberFormatException, IOException {
		System.out.println("Binary search algorithm for arrays.");
		System.out.println("Enter the number of students registered for the class:");
		InputStreamReader inputStreamReader = new InputStreamReader(System.in);
		BufferedReader br = new BufferedReader(inputStreamReader);
		
		int numberOfStudents = Integer.parseInt(br.readLine());
		OrderedAttendances attendance = new OrderedAttendances(numberOfStudents);
		attendance.addStudent(55);
		attendance.addStudent(50);
		attendance.addStudent(52);
		attendance.addStudent(5);
		attendance.addStudent(3);
		attendance.addStudent(4);
		attendance.addStudent(2);
		attendance.addStudent(1);
		attendance.displayCurrentAttendance();
		attendance.searchStudent(52);
		attendance.searchStudent(10);
	}
}
