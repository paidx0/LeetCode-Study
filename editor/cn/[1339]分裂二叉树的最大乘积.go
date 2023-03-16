// 给你一棵二叉树，它的根为 root 。请你删除 1 条边，使二叉树分裂成两棵子树，且它们子树和的乘积尽可能大。
//
// 由于答案可能会很大，请你将结果对 10^9 + 7 取模后再返回。
//
//
//
// 示例 1：
//
//
//
// 输入：root = [1,2,3,4,5,6]
// 输出：110
// 解释：删除红色的边，得到 2 棵子树，和分别为 11 和 10 。它们的乘积是 110 （11*10）
//
//
// 示例 2：
//
//
//
// 输入：root = [1,null,2,3,4,null,null,5,6]
// 输出：90
// 解释：移除红色的边，得到 2 棵子树，和分别是 15 和 6 。它们的乘积为 90 （15*6）
//
//
// 示例 3：
//
// 输入：root = [2,3,9,10,7,8,6,5,4,11,1]
// 输出：1025
//
//
// 示例 4：
//
// 输入：root = [1,1]
// 输出：1
//
//
//
//
// 提示：
//
//
// 每棵树最多有 50000 个节点，且至少有 2 个节点。
// 每个节点的值在 [1, 10000] 之间。
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 86 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// func maxProduct(root *TreeNode) int {
// 	var res int
// 	rootSum := sumdfs(root)
// 	dfs(root, &res, rootSum)
//
// 	return res % (1e9 + 7)
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
// func sumdfs(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	return root.Val + sumdfs(root.Left) + sumdfs(root.Right)
// }
//
// // 后序遍历，边遍历边算切分以后得子树和乘积
// func dfs(root *TreeNode, res *int, rootsum int) int {
// 	if root == nil {
// 		return 0
// 	}
// 	left := dfs(root.Left, res, rootsum)
// 	right := dfs(root.Right, res, rootsum)
// 	cursum := root.Val + left + right
//
// 	*res = max(*res, (rootsum-cursum)*cursum)
// 	return cursum
// }

// 求出总和，和平分时乘积是最大的
// avg是和的一半，half是真正的子树一半接近avg的
func maxProduct(root *TreeNode) int {
	half, avg := 0, 0
	rootsum := total(root, &half, avg)
	half, avg = rootsum, rootsum/2
	total(root, &half, avg)
	return (rootsum - half) * half % (1e9 + 7)
}
func total(root *TreeNode, half *int, avg int) int {
	if root == nil {
		return 0
	}
	cursum := root.Val + total(root.Left, half, avg) + total(root.Right, half, avg)
	if abs(cursum-avg) < abs(*half-avg) {
		*half = cursum
	}
	return cursum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// leetcode submit region end(Prohibit modification and deletion)
