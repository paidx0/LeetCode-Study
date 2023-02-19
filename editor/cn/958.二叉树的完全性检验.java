/*
 * @lc app=leetcode.cn id=958 lang=java
 *
 * [958] 二叉树的完全性检验
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
    int maxIndex = 0;
    int nullIndex = 1001;
    boolean ans = true;
    public boolean isCompleteTree(TreeNode root) {
        DFS(root, 0, 1);
        return ans;
    }
    public void DFS(TreeNode root, int depth, int index){
        if(root == null || ans == false){
            nullIndex = Math.min(nullIndex, index);
            if(nullIndex < maxIndex) ans = false;
            return;
        }
        if(index > maxIndex)
            maxIndex = index;
        DFS(root.right, depth + 1, index * 2 + 1);
        DFS(root.left, depth + 1, index * 2);
        
    }
}
// @lc code=end

