/*
 * @lc app=leetcode.cn id=98 lang=java
 *
 * [98] 验证二叉搜索树
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
    boolean ans = true;
    Integer pre = null;
    public boolean isValidBST(TreeNode root) {
        inorder(root);
        return ans;
    }
    public void inorder(TreeNode root){
        if(ans == false) return;
        if(root == null) return;
        isValidBST(root.left);
        if(pre == null){
            pre = root.val;
        }else if(root.val <= pre){
            ans = false;
        }
        pre = root.val;
        isValidBST(root.right);
    }
}
// @lc code=end

