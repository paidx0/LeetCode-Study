/*
 * @lc app=leetcode.cn id=662 lang=java
 *
 * [662] 二叉树最大宽度
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
    Map<Integer, Integer> res = new HashMap<>();
    int maxWidth = 1;
    public int widthOfBinaryTree(TreeNode root) {
        DFS(root, 1, 0);
        return maxWidth;
    }
    public void DFS(TreeNode root, int index, int depth){
        if(root == null) return;
        if(res.size() == depth){
            res.put(depth, index);
        }else if(index - res.get(depth) + 1 > maxWidth){
            maxWidth = index - res.get(depth) + 1;
        }
        DFS(root.left, index * 2, depth + 1);
        DFS(root.right, index * 2 + 1, depth + 1);
    }
}
// @lc code=end

