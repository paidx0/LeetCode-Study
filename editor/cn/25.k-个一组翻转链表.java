import Test.ListNode;

/*
 * @lc app=leetcode.cn id=25 lang=java
 *
 * [25] K 个一组翻转链表
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
    public ListNode reverseKGroup(ListNode head, int k) {
        ListNode tmp = new ListNode();
        tmp.next = head;
        visit(tmp, k);
        return tmp.next;
    }
    public void visit(ListNode head, int k){
        frt = head;
        if(!reverse(frt.next, k)) return;
        frt.next = nxt;
        visit(frt, k);
    }
    public boolean reverse(ListNode head, int k){
        if(k > 0 && head == null) return false;
        if(k == 0){
            nxt = head;
            return true;
        }
        if(reverse(head.next, k - 1)){
            frt.next = head;
            frt = head;
            return true;
        }
        return false;
    }
}
// @lc code=end

