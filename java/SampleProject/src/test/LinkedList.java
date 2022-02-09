package test;

// Data structure of a node.
class Node {
	// To store data.
	int data;
	// To store the link to next node. Basically the next node itself.
	Node node;
}


public class LinkedList {
	// Used to locate the start of linked list.
	Node head = null;
	// Used to located the end of linked list.
	Node last = null;
	
	// Add element at the last position.
	void addAtLast(int data) {
		// Create the node.
		Node current = new Node();
		current.data = data;
		// Check if the linked list is empty.
		if(last == null) {
			head = last = current;
		} else { // Add to the last node.
			last.node = current;
			last = current;
		}
	}
	
	// Add element at the first position.
	void addAtFirst(int data) {
		// Create the node.
		Node current = new Node();
		current.data = data;
		// Check if the linked list is empty. 
		if(head == null) {
			head = last = current;
		} else {
			// Add to the first node.
			current.node = head;
			head = current;
		}
	}
	
	// To show all the elements.
	void display() {
		// Start with the first node.
		Node current = head;
		
		// Loop till the node is null.
		while(current != null) {
			System.out.println(current.data);
			// Move to the next node.
			current = current.node;
		}
	}
	
	void remove(int data) {
		Node current = head;
		Node previous = current;
		int count = 0;
		while(current != null) {
			if(current.data == data) {
				// Extreme condition: First node.
				if(previous == current) {
					head = last = current.node;
					current.node = null;
				} else {
					previous.node = current.node;
					current.node = null;
				}
				System.out.println("Removed data at index: " + count);
				break;
			}
			
			previous = current;
			current = current.node;
			count++;
			
		}
	}
	
	void search(int data) {
		// Start with the first node.
		Node current = head;
		int count = 0;
		// Loop till the node is null.
		while(current != null) {
			if(current.data == data) {
				System.out.println("Searched element is present at index: " + count);
			}
			// Move to the next node.
			current = current.node;
			count++;
		} 
	}
	
	public static void main(Strings[] args) {
		LinkedList list = new LinkedList();
		list.addAtLast(1);
		list.display();
		list.remove(1);
		list.display();
		list.addAtLast(2);
		list.addAtLast(3);
		list.addAtFirst(4);
		list.display();
		list.search(3);
		list.remove(3);
		list.display();
	}
}
