/*
 * @lc app=leetcode.cn id=572 lang=java
 *
 * [572] 另一棵树的子树
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
    int lx = 56789;
    int mx = 45678;
    int rx = 34567;
    int subHash;
    int flag = 0;
    public boolean isSubtree(TreeNode root, TreeNode subRoot) {
        subHash = trans1(subRoot);
        trans2(root);
        return flag == 1;
    }
    public int trans1(TreeNode root){
        if(root == null) return 1;
        int hash = root.val + mx + trans1(root.left) * lx + trans1(root.right) * rx;
        root.val = hash;
        return hash;
    }
    public int trans2(TreeNode root){
        if(root == null) return 1;
        int hash = root.val + mx + trans2(root.left) * lx + trans2(root.right) * rx;
        root.val = hash;
        if(hash == subHash) flag = 1;
        return hash;
    }
}
// @lc code=end

