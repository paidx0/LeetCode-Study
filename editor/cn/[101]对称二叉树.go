//给你一个二叉树的根节点 root ， 检查它是否轴对称。
//
//
//
// 示例 1：
//
//
//输入：root = [1,2,2,3,4,4,3]
//输出：true
//
//
// 示例 2：
//
//
//输入：root = [1,2,2,null,3,null,3]
//输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 1000] 内
// -100 <= Node.val <= 100
//
//
//
//
// 进阶：你可以运用递归和迭代两种方法解决这个问题吗？
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 2275 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	return root == nil || dfs(root.Left, root.Right)
}
func dfs(rootLeft, rootRight *TreeNode) bool {
	if rootLeft == nil && rootRight == nil {
		return true
	}
	if rootLeft == nil || rootRight == nil {
		return false
	}
	if rootLeft.Val != rootRight.Val {
		return false
	}
	return dfs(rootLeft.Left, rootRight.Right) && dfs(rootLeft.Right, rootRight.Left)
}

//leetcode submit region end(Prohibit modification and deletion)
