/*
 * @lc app=leetcode.cn id=429 lang=java
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
    List<List<Integer>> ans = new ArrayList();
    public List<List<Integer>> levelOrder(Node root) {
        DFS(root, 0);
        return ans;
    }
    public void DFS(Node root, int depth){
        if(root == null) return;
        if(depth == ans.size()) ans.add(new ArrayList<Integer>());
        List<Integer> tmp = ans.get(depth);
        tmp.add(root.val);
        root.children.forEach(node -> {
            DFS(node, depth + 1);
        });
    }
}
// @lc code=end

