import java.util.ArrayList;

/*
 * @lc app=leetcode.cn id=95 lang=java
 *
 * [95] 不同的二叉搜索树 II
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
    public List<TreeNode> generateTrees(int n) {
        List<TreeNode> lst = new ArrayList<>();
        if(n == 0) return lst;
        if(n == 1) {
            TreeNode node = new TreeNode(1);
            lst.add(node);
            return lst;
        }
        List<TreeNode> res = generateTrees(n - 1);
        for(int i = 0; i < res.size(); i++){
            TreeNode tmp = res.get(i);
            TreeNode node = new TreeNode(n);
            TreeNode resNode = treeCopy(tmp);
            node.left = resNode;
            lst.add(node);
            TreeNode rightKid = tmp;
            while(rightKid != null){
                resNode = treeCopy(tmp);
                node = new TreeNode(n);
                lst.add(resNode);
                while(resNode.val != rightKid.val) resNode = resNode.right;
                node.left = resNode.right;
                resNode.right = node;
                rightKid = rightKid.right;
            }

        }
        return lst;
    }
    public TreeNode treeCopy(TreeNode root){
        if(root == null) return null;
        TreeNode node = new TreeNode();
        node.val = root.val;
        node.left = treeCopy(root.left);
        node.right = treeCopy(root.right);
        return node;
    }
}
// @lc code=end

