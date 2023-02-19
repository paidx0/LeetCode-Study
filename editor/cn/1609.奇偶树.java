import java.util.ArrayList;
import java.util.List;

/*
 * @lc app=leetcode.cn id=1609 lang=java
 *
 * [1609] 奇偶树
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
    List<Integer> res = new ArrayList<>();
    public boolean isEvenOddTree(TreeNode root) {
        DFS(root, 0);
        return ans;
    }
    public void DFS(TreeNode root, int depth){
        if(root == null || ans == false) return;
        if(res.size() == depth) {
            if(depth % 2 == 0 && root.val % 2 == 0){
                ans = false;
                return;
            }
            if(depth % 2 != 0 && root.val % 2 != 0){
                ans = false;
                return;
            }
            res.add(root.val);
        }
        else if(depth % 2 == 0){
            if(root.val % 2 == 0) {
                ans = false;
                return;
            }
            else if(res.get(depth) >= root.val) {
                ans = false;
                return;
            }
            else res.set(depth, root.val);
        }
        else{
            if(root.val % 2 != 0) {
                ans = false;
                return;
            }
            else if(res.get(depth) <= root.val) {
                ans = false;
                return;
            }
            else res.set(depth, root.val);
        }
        DFS(root.left, depth + 1);
        DFS(root.right, depth + 1);
    }
}
// @lc code=end

