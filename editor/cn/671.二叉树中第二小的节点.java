/*
 * @lc app=leetcode.cn id=671 lang=java
 *
 * [671] 二叉树中第二小的节点
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    int ans;
    public int findSecondMinimumValue(TreeNode root) {
        if(root == null || root.left == null) return -1;
        int left, right;
        left = root.left.val == root.val ?
            findSecondMinimumValue(root.left) :
            root.left.val;
        right = root.right.val == root.val ?
            findSecondMinimumValue(root.right) :
            root.right.val;
        int min = Math.min(left, right);
        return min > 0 ? min : Math.max(left, right);
    }
}
// @lc code=end

