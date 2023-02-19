import java.util.Comparator;
import java.util.List;

/*
 * @lc app=leetcode.cn id=1305 lang=java
 *
 * [1305] 两棵二叉搜索树中的所有元素
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
    List<Integer> ansList = new ArrayList();
    public List<Integer> getAllElements(TreeNode root1, TreeNode root2) {
        iot(root1);
        iot(root2);
        Comparator c = new Comparator<Integer>(){
            @Override
            public int compare(Integer a, Integer b){
                return a - b;
            }
        };
        ansList.sort(c);
        return ansList;
    }
    public void iot(TreeNode root){
        if(root == null) return;
        iot(root.left);
        ansList.add(root.val);
        iot(root.right);
    }
}
// @lc code=end

