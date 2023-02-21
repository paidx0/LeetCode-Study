/*
 * @lc app=leetcode.cn id=234 lang=java
 *
 * [234] 回文链表
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
    ListNode node;
    public boolean isPalindrome(ListNode head) {
        node = head;
        return dfs(head);
    }
    public boolean dfs(ListNode head){
        boolean res = true;
        if(head.next != null) res = dfs(head.next);
        if(res == false) return false;
        if(node.val != head.val) return false;
        else{
            node = node.next;
            return true;
        }
    }
}
// @lc code=end

