//给定一个二叉搜索树 root 和一个目标结果 k，如果二叉搜索树中存在两个元素且它们的和等于给定的目标结果，则返回 true。
//
//
//
// 示例 1：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 9
//输出: true
//
//
// 示例 2：
//
//
//输入: root = [5,3,6,2,4,null,7], k = 28
//输出: false
//
//
//
//
// 提示:
//
//
// 二叉树的节点个数的范围是 [1, 10⁴].
// -10⁴ <= Node.val <= 10⁴
// 题目数据保证，输入的 root 是一棵 有效 的二叉搜索树
// -10⁵ <= k <= 10⁵
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 哈希表 双指针 二叉树 👍 453 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Map存访问过的值，边遍历边查map就行了
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]int, 0)
	return dfs(root, k, m)
}
func dfs(root *TreeNode, k int, m map[int]int) bool {
	if root == nil {
		return false
	}
	if _, ok := m[k-root.Val]; ok {
		return true
	}
	m[root.Val]++
	return dfs(root.Left, k, m) || dfs(root.Right, k, m)
}

//leetcode submit region end(Prohibit modification and deletion)
