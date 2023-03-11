// 给定一个二叉树的
// root ，返回 最长的路径的长度 ，这个路径中的 每个节点具有相同值 。 这条路径可以经过也可以不经过根节点。
//
// 两个节点之间的路径长度 由它们之间的边数表示。
//
//
//
// 示例 1:
//
//
//
//
// 输入：root = [5,4,5,1,1,5]
// 输出：2
//
//
// 示例 2:
//
//
//
//
// 输入：root = [1,4,5,4,4,5]
// 输出：2
//
//
//
//
// 提示:
//
//
// 树的节点数的范围是
// [0, 10⁴]
// -1000 <= Node.val <= 1000
// 树的深度将不超过 1000
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 756 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	var dfs func(*TreeNode) int
	var ans int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		left1, right1 := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			left1 = left + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			right1 = right + 1
		}
		ans = max(ans, left1+right1)
		return max(left1, right1)
	}
	dfs(root)
	return ans
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// leetcode submit region end(Prohibit modification and deletion)
