/*
 * @lc app=leetcode.cn id=1008 lang=java
 *
 * [1008] 前序遍历构造二叉搜索树
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
    public TreeNode bstFromPreorder(int[] preorder) {
        return build(preorder, 0, preorder.length - 1);
    }
    public TreeNode build(int[] preorder, int start, int end){
        if(start > end) return null;
        int p = start;
        TreeNode root = new TreeNode(preorder[p++]);
        while(p <= end && preorder[p] < root.val) p++;
        root.left = build(preorder, start + 1, p - 1);
        root.right = build(preorder, p, end);
        return root;
    }
}
// @lc code=end

