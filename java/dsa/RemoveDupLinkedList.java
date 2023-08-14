class Solution {
    public ListNode deleteDuplicates(ListNode head) {
        if(head == null || head.next == null) {
            return head;
        }
        ListNode curr = head.next;
        ListNode prev = head;
        
        while(curr.next != null) {
            if(prev.val != curr.val) {
                prev.next = curr;
                prev = prev.next;
            }
            curr = curr.next;
        }
        
        if(prev.val == curr.val) {
            prev.next = curr.next;
        } else {
            prev.next = curr;
        }
        
        return head;
    }
}