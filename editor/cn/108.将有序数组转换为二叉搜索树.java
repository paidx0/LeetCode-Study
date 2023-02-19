/*
 * @lc app=leetcode.cn id=108 lang=java
 *
 * [108] 将有序数组转换为二叉搜索树
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
    public TreeNode sortedArrayToBST(int[] nums) {
        if(nums.length == 0) return null;
        int len = nums.length;
        TreeNode root = new TreeNode(nums[len / 2]);
        root.left = sortedArrayToBST(Arrays.copyOfRange(nums, 0, len / 2));
        root.right = sortedArrayToBST(Arrays.copyOfRange(nums,len / 2 + 1, len));
        return root;
    }
}
// @lc code=end

