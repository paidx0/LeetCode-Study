/*
 * @lc app=leetcode.cn id=206 lang=java
 *
 * [206] 反转链表
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
    ListNode newList;
    ListNode temp;
    public ListNode reverseList(ListNode head) {
        newList = null;
        temp = null;
        dfs(head);
        return newList;
    }
    public void dfs(ListNode head){
        if(head == null) return;
        dfs(head.next);
        if(newList == null){
            newList = head;
            temp = head;
        }else{
            temp.next = head;
            head.next = null;
            temp = head;
        }
    }
}
// @lc code=end

