import java.util.List;

/*
 * @lc app=leetcode.cn id=1161 lang=java
 *
 * [1161] 最大层内元素和
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
    List<Integer> map = new ArrayList();
    int max;
    int ans = 1;
    public int maxLevelSum(TreeNode root) {
        DFS(root, 0);
        max = map.get(0);
        for(int i = 1; i <= map.size(); i++){
            if(map.get(i - 1) > max){
                max = map.get(i - 1);
                ans = i;
            }
        }
        return ans;
    }
    public void DFS(TreeNode root, int depth){
        if(root == null) return;
        if(depth == map.size()){
            map.add(root.val);
        }
        else{
            int tmp = map.get(depth) + root.val;
            map.set(depth, tmp);
        }
        DFS(root.left, depth + 1);
        DFS(root.right, depth + 1);
    }
}
// @lc code=end

