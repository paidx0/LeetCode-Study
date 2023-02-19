/*
 * @lc app=leetcode.cn id=919 lang=java
 *
 * [919] 完全二叉树插入器
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
class CBTInserter {
    TreeNode rootNode;
    int maxDepth = -1;
    TreeNode insertNode;
    boolean insertLeft;
    public CBTInserter(TreeNode root) {
        rootNode = root;
        getDepth(root);
    }
    
    public int insert(int val) {
        insertNode = null;
        insertDfs(rootNode, 0);
        if(insertNode == null){
            maxDepth += 1;
            insertDfs(rootNode, 0);
        }
        TreeNode node = new TreeNode(val);
        if(insertLeft) insertNode.left = node;
        else insertNode.right = node;
        return insertNode.val;
    }
    
    public void insertDfs(TreeNode root, int depth){
        if(root == null || insertNode != null || depth == maxDepth) return;
        if(depth == maxDepth - 1){
            if(root.left == null){
                insertNode = root;
                insertLeft = true;
            }else if(root.right == null){
                insertNode = root;
                insertLeft = false;
            }
        }
        insertDfs(root.left, depth + 1);
        insertDfs(root.right, depth + 1);
    }
    public TreeNode get_root() {
        return rootNode;
    }

    public void getDepth(TreeNode root){
        if(root == null) return;
        maxDepth += 1;
        getDepth(root.left);
    }
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * CBTInserter obj = new CBTInserter(root);
 * int param_1 = obj.insert(val);
 * TreeNode param_2 = obj.get_root();
 */
// @lc code=end

