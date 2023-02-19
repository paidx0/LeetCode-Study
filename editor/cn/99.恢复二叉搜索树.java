/*
 * @lc app=leetcode.cn id=99 lang=java
 *
 * [99] 恢复二叉搜索树
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
    TreeNode pre = null;
    TreeNode x = null;
    TreeNode y = null;
    boolean flag = false;
    public void recoverTree(TreeNode root) {
        inorder(root);
        x.val ^= y.val;
        y.val ^= x.val;
        x.val ^= y.val;
    }
    public void inorder(TreeNode root){
        if(root == null) return;
        inorder(root.left);
        if(pre == null) pre = root;
        else{
            if(root.val < pre.val && !flag){
                x = pre;
                y = root;
                flag = true;
            }
            if(root.val < pre.val && flag){
                y = root;
            }
        }
        pre = root;
        inorder(root.right);
    }
}
// @lc code=end

