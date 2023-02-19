/*
 * @lc app=leetcode.cn id=623 lang=java
 *
 * [623] 在二叉树中增加一行
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
    public TreeNode addOneRow(TreeNode root, int val, int depth) {
        if(depth == 1){
            TreeNode node = new TreeNode(val, root, null);
            return node;
        }
        DFS(root, 1, val, depth);
        return root;
    }
    public void DFS(TreeNode root, int depth, int val, int target){
        if(root == null) return;
        if(depth == target - 1){
            TreeNode tmp1 = new TreeNode(val, root.left, null);
            TreeNode tmp2 = new TreeNode(val, null, root.right);
            root.left = tmp1;
            root.right = tmp2;
            return;
        }
        DFS(root.left, depth + 1, val, target);
        DFS(root.right, depth + 1, val, target);
    }
}
// @lc code=end

