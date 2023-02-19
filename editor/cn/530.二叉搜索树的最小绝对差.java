/*
 * @lc app=leetcode.cn id=530 lang=java
 *
 * [530] 二叉搜索树的最小绝对差
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
    int diff = 100000;
    int preVal = -1;
    public int getMinimumDifference(TreeNode root) {
        getDiff(root);
        return diff;
    }
    public void getDiff(TreeNode root){
        if(root == null) return;
        if(root.left != null) getDiff(root.left);
        if(preVal != -1) diff = Math.min(root.val - preVal, diff);
        preVal = root.val;
        if(root.right != null) getDiff(root.right);
    }
}
// @lc code=end

