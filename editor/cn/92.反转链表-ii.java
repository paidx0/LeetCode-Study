import Test.ListNode;

/*
 * @lc app=leetcode.cn id=92 lang=java
 *
 * [92] 反转链表 II
 */

// @lc code=start
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
    ListNode frt;
    ListNode nxt;
    public ListNode reverseBetween(ListNode head, int left, int right) {
        int i = 0;
        ListNode pre = new ListNode(0, head);
        head = pre;
        while(i++ < left - 1) head = head.next;
        frt = head;
        reverse(head.next, right - left);
        frt.next = nxt;
        return pre.next;
    }
    public void reverse(ListNode head, int k){
        if(k < 0|| head == null) return;
        if(k == 0){
            nxt = head.next;
        }
        reverse(head.next, k - 1);
        frt.next = head;
        frt = head;
    }
}
// @lc code=end

