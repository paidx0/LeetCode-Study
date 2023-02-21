/*
 * @lc app=leetcode.cn id=24 lang=java
 *
 * [24] 两两交换链表中的节点
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
    public ListNode swapPairs(ListNode head) {
        ListNode tmp = new ListNode();
        tmp.next = head;
        dfs(tmp);
        return tmp.next;
    }
    public void dfs(ListNode head){
        ListNode tmp1 = head.next;
        if(tmp1 == null) return;
        ListNode tmp2 = tmp1.next;
        if(tmp2 == null) return;
        head.next = tmp2;
        tmp1.next = tmp2.next;
        tmp2.next = tmp1;
        dfs(tmp1);
    }
}
// @lc code=end

