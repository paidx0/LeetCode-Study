import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=230 lang=java
 *
 * [230] 二叉搜索树中第K小的元素
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
    int count = 0;
    int ans;
    public int kthSmallest(TreeNode root, int k) {
        inorder(root, k);
        return ans;
    }
    public void inorder(TreeNode root, int k){
        if(count >= k) return;
        if(root == null) return;
        inorder(root.left, k);
        count++;
        if(count == k) ans = root.val;
        inorder(root.right, k);
    }
}
// @lc code=end

