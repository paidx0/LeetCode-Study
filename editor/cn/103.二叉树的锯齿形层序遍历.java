import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.Queue;

/*
 * @lc app=leetcode.cn id=103 lang=java
 *
 * [103] 二叉树的锯齿形层序遍历
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
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        if(root == null) return new ArrayList<>();
        boolean flag = true;
        int logSize = 1;
        Deque<TreeNode> que = new ArrayDeque<>();
        List<List<Integer>> ans = new ArrayList<>();
        que.offer(root);
        List<Integer> temp = new ArrayList();
        while(que.size() != 0){
            if(flag){
                TreeNode node = que.pollFirst();
                temp.add(node.val);
                if(node.left != null)
                    que.offerLast(node.left);
                if(node.right != null)
                    que.offerLast(node.right);
                logSize -= 1;
            }else{
                TreeNode node = que.pollLast();
                temp.add(node.val);
                if(node.right != null)
                    que.offerFirst(node.right);
                if(node.left != null)
                    que.offerFirst(node.left);
                logSize -= 1;
            }
            if(logSize == 0){
                flag = !flag;
                logSize = que.size();
                ans.add(temp);
                temp = new ArrayList();
            }
        }
        return ans;
    }
}
// @lc code=end

