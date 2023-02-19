import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Queue;

/*
 * @lc app=leetcode.cn id=501 lang=java
 *
 * [501] 二叉搜索树中的众数
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
    int base = 100001;
    int maxCount = 0;
    List<Integer> res = new ArrayList<>();
    public int[] findMode(TreeNode root) {
        inorder(root);
        int[] ans = new int[res.size()];
        for(int i = 0; i < res.size(); i++){
            ans[i] = res.get(i);
        }
        return ans;
    }
    public void inorder(TreeNode root){
        if(root == null) return;
        inorder(root.left);
        if(root.val == base) count += 1;
        else{
            count = 1;
            base = root.val;
        }
        if(count > maxCount){
            res.clear();
            res.add(base);
            maxCount = count;
        }else if (count == maxCount){
            res.add(base);
        }
        inorder(root.right);
    }
}
// @lc code=end

