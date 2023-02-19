/*
 * @lc app=leetcode.cn id=117 lang=java
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
    public int val;
    public Node left;
    public Node right;
    public Node next;

    public Node() {}
    
    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, Node _left, Node _right, Node _next) {
        val = _val;
        left = _left;
        right = _right;
        next = _next;
    }
};
*/

class Solution {
    public Node connect(Node root) {
        if(root == null) return root;
        if(root.right != null){
            if(root.left != null)
                root.left.next = root.right;
        }
        Node node = getNode(root.next);
        if(node != null){
            if(root.right != null){
                if(node.left != null){
                    root.right.next = node.left;
                }else if(node.right != null){
                    root.right.next = node.right;
                }
            }else if(root.left != null){
                if(node.left != null){
                    root.left.next = node.left;
                }else if(node.right != null){
                    root.left.next = node.right;
                }
            }
        }
        connect(root.right);
        connect(root.left);
        return root;
    }
    public Node getNode(Node root){
        if(root == null) return null;
        if(root.left != null || root.right != null) return root;
        return getNode(root.next);
    }
}
// @lc code=end

