//给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。
//
// 差值是一个正数，其数值等于两值之差的绝对值。
//
//
//
// 示例 1：
//
//
//输入：root = [4,2,6,1,3]
//输出：1
//
//
// 示例 2：
//
//
//输入：root = [1,0,48,null,null,12,49]
//输出：1
//
//
//
//
// 提示：
//
//
// 树中节点的数目范围是 [2, 10⁴]
// 0 <= Node.val <= 10⁵
//
//
//
//
// 注意：本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-
//nodes/ 相同
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉搜索树 二叉树 👍 429 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 中序遍历，前一个和后一个比较
func getMinimumDifference(root *TreeNode) int {
	res, pre := math.MaxInt16, -1
	dfs(root, &res, &pre)
	return res
}
func dfs(root *TreeNode, res, pre *int) {
	if root == nil {
		return
	}
	dfs(root.Left, res, pre)

	if *pre != -1 {
		*res = min(*res, abs(root.Val-*pre))
	}
	*pre = root.Val

	dfs(root.Right, res, pre)
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

//leetcode submit region end(Prohibit modification and deletion)
