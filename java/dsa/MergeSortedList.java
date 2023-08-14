/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {


// Recursion
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        if(list1 == null) return list2;
        else if(list2 == null) return list1;
        boolean isOneBigger = list1.val > list2.val;
        ListNode current = isOneBigger?list2:list1;
        current.next = isOneBigger?mergeTwoLists(list1,list2.next):mergeTwoLists(list1.next,list2);
        return current;
    }


// Without creating DS
    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {

    // Creating a new linked list to accomodate all the elements of the merged linked list
    ListNode node = new ListNode(0);    // it will assign the value of node as 0
    ListNode head = node;
    
    while(list1 != null && list2 != null)
    {
        if(list1.val <= list2.val)
        {
            node.next = list1;
            list1 = list1.next;
            node = node.next;
        }
        else
        {
            node.next = list2;
            list2 = list2.next;
            node = node.next;
        }
    }
    
    // if our list1 got null, then add the rest of the list2 elements in the node
    if(list1 == null)
    {
        node.next = list2;
    }
    else
    {
        node.next = list1;
    }
    
    return head.next;
}


    public ListNode mergeTwoLists(ListNode list1, ListNode list2) {
        ListNode list3 = null, head = null;
        
        while(list1 != null && list2 != null) {
            
            if(head == null) {
                list3 = new ListNode();
                head = list3;
            } else {
                list3.next = new ListNode();
                list3 = list3.next;
            }
            if(list1.val <= list2.val) {
                list3.val = list1.val;
                list1 = list1.next;
            } else {
                list3.val = list2.val;
                list2 = list2.next;
            }
        }
        
         while(list1 != null) {
             if(head == null) {
                list3 = new ListNode();
                head = list3;
            } else {
             list3.next = new ListNode();
            list3 = list3.next;
             }
             list3.val = list1.val;
              list1 = list1.next;
             
         }
        while(list2 != null) {
            if(head == null) {
                list3 = new ListNode();
                head = list3;
            } else {
             list3.next = new ListNode();
            list3 = list3.next;
             }
             list3.val = list2.val;
              list2 = list2.next;
             
         }
        return head;
    }
}