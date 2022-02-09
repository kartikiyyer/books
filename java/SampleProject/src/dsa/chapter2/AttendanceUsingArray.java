package dsa.chapter2;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

class Attendance {
	private int[] students;
	private int numberOfStudentsPresent;
	public Attendance(int numberOfStudents) {
		numberOfStudentsPresent = 0;
		students = new int[numberOfStudents]; 
	}
	
	public void displayCurrentAttendance() {
		System.out.println("Students present in the class:");
		for(int i=0; i<numberOfStudentsPresent; i++) {
			System.out.println(students[i]);
		}
	}
	
	public void addStudent(int rollNumber) {
		System.out.println("Adding student with roll number: " + rollNumber);
		students[numberOfStudentsPresent++] = rollNumber;
	}
	
	public void searchStudent(int rollNumber) {
		int i;
		for(i=0; i<numberOfStudentsPresent; i++) {
			if(students[i] == rollNumber) {
				System.out.println("Student found with roll number: " + rollNumber);
				break;
			}
		}
		
		if(i == numberOfStudentsPresent) {
			System.out.println("Student with roll number: " + rollNumber + " is not present.");
		}
	}
	
	public void deleteStudent(int rollNumber) {
		System.out.println("Deleting student with roll number: " + rollNumber);
		int i;
		for(i=0; i<numberOfStudentsPresent; i++) {
			if(students[i] == rollNumber) {
				break;
			}
		}
		
		if(i == numberOfStudentsPresent) {
			System.out.println("Student with roll number: " + rollNumber + " is not present.");
		} else {
			while(i < numberOfStudentsPresent-1) {
				students[i] = students[i+1];
				i++;
			}
		}
		numberOfStudentsPresent--;
	}
}


public class AttendanceUsingArray {
	public static void main(String args[]) throws NumberFormatException, IOException {
		System.out.println("This is the implementation of an Attendance application for today...");
		System.out.println("Enter the number of students registered for the class:");
		InputStreamReader inputStreamReader = new InputStreamReader(System.in);
		BufferedReader br = new BufferedReader(inputStreamReader);
		
		int numberOfStudents = Integer.parseInt(br.readLine());
		Attendance attendance = new Attendance(numberOfStudents);
		attendance.addStudent(55);
		attendance.addStudent(50);
		attendance.addStudent(52);
		attendance.addStudent(1);
		attendance.displayCurrentAttendance();
		attendance.searchStudent(52);
		attendance.searchStudent(10);
		attendance.deleteStudent(52);
		attendance.deleteStudent(50);
		attendance.deleteStudent(10);
		attendance.displayCurrentAttendance();
		
	}
}

