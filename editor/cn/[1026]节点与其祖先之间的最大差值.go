// 给定二叉树的根节点 root，找出存在于 不同 节点 A 和 B 之间的最大值 V，其中 V = |A.val - B.val|，且 A 是 B 的祖先。
//
//
// （如果 A 的任何子节点之一为 B，或者 A 的任何子节点是 B 的祖先，那么我们认为 A 是 B 的祖先）
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [8,3,10,1,6,null,14,null,null,4,7,13]
// 输出：7
// 解释：
// 我们有大量的节点与其祖先的差值，其中一些如下：
// |8 - 3| = 5
// |3 - 7| = 4
// |8 - 1| = 7
// |10 - 13| = 3
// 在所有可能的差值中，最大值 7 由 |8 - 1| = 7 得出。
//
//
// 示例 2：
//
//
// 输入：root = [1,null,2,null,0,3]
// 输出：3
//
//
//
//
// 提示：
//
//
// 树中的节点数在 2 到 5000 之间。
// 0 <= Node.val <= 10⁵
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 137 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 深搜即可，root与左右子树最大和最小值比较，返回差值最大的绝对值
func maxAncestorDiff(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}
func dfs(root *TreeNode, res *int) (int, int) {
	if root == nil {
		return -1, -1
	}
	leftMax, leftMin := dfs(root.Left, res)
	if leftMax == -1 && leftMin == -1 {
		leftMax = root.Val
		leftMin = root.Val
	}
	rightMax, rightMin := dfs(root.Right, res)
	if rightMax == -1 && rightMin == -1 {
		rightMax = root.Val
		rightMin = root.Val
	}
	*res = max(*res, max(abs(root.Val-min(leftMin, rightMin)), abs(root.Val-max(leftMax, rightMax))))
	return max(leftMax, max(rightMax, root.Val)), min(leftMin, min(rightMin, root.Val))
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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

// leetcode submit region end(Prohibit modification and deletion)
