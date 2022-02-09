import java.util.*;

class Node {
	int data;
	Node next;
}
public class LinkedListReverse {
	public static void main(String[] args) {
		Node node = new Node();
		node.data = 1;
		node.next = new Node();
		node.next.data = 2;
		node.next.next = new Node();
		node.next.next.data = 3;

		// Using stack.
		Stack<Integer> s = new Stack<Integer>();
		// Using string and printing in reverse.
		String str = new String();
		Node head = node;
		while(head != null) {
			System.out.println(head.data);
			s.push(head.data);
			str += head.data;
			head = head.next;
		} 

		while(s.size() != 0) {
			System.out.println(s.pop());
		}


		for(int i=str.length()-1; i >=0; i--) {
			System.out.println(str.charAt(i));
		}


	}
}

