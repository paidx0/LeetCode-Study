/*
 * @lc app=leetcode.cn id=1302 lang=java
 *
 * [1302] 层数最深叶子节点的和
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
    int maxDepth = 0;
    int ans = 0;
    public int deepestLeavesSum(TreeNode root) {
        DFS(root, 0);
        return ans;
    }
    public void DFS(TreeNode root, int depth){
        if(root == null) return;
        if(depth == maxDepth) ans += root.val;
        if(depth > maxDepth){
            maxDepth = depth;
            ans = root.val;
        }
        DFS(root.left, depth + 1);
        DFS(root.right, depth + 1);
    }
}
// @lc code=end

