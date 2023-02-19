/*
 * @lc app=leetcode.cn id=107 lang=java
 *
 * [107] 二叉树的层序遍历 II
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
    List<List<Integer>> ans = new ArrayList();
    public List<List<Integer>> levelOrderBottom(TreeNode root) {
        DFS(root, 0);
        return ans;
    }
    public void DFS(TreeNode root, int depth){
        if(root == null) return;
        if(depth == ans.size()) ans.add(0, new ArrayList<Integer>());
        List<Integer> tmp = ans.get(ans.size() - 1 - depth);
        tmp.add(root.val);
        DFS(root.left, depth + 1);
        DFS(root.right, depth + 1);
    }
}
// @lc code=end

