// 给你一棵二叉树，请你返回满足以下条件的所有节点的值之和：
//
//
// 该节点的祖父节点的值为偶数。（一个节点的祖父节点是指该节点的父节点的父节点。）
//
//
// 如果不存在祖父节点值为偶数的节点，那么返回 0 。
//
//
//
// 示例：
//
//
//
// 输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
// 输出：18
// 解释：图中红色节点的祖父节点的值为偶数，蓝色节点为这些红色节点的祖父节点。
//
//
//
//
// 提示：
//
//
// 树中节点的数目在 1 到 10^4 之间。
// 每个节点的值在 1 到 100 之间。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 84 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 深搜是偶数就加上孙子节点的值
func sumEvenGrandparent(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Val%2 == 0 {
		if root.Left != nil {
			if root.Left.Left != nil {
				sum += root.Left.Left.Val
			}
			if root.Left.Right != nil {
				sum += root.Left.Right.Val
			}
		}
		if root.Right != nil {
			if root.Right.Left != nil {
				sum += root.Right.Left.Val
			}
			if root.Right.Right != nil {
				sum += root.Right.Right.Val
			}
		}
	}
	return sum + sumEvenGrandparent(root.Left) + sumEvenGrandparent(root.Right)
}

// leetcode submit region end(Prohibit modification and deletion)
