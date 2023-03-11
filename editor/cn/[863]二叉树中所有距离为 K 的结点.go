// 给定一个二叉树（具有根结点 root）， 一个目标结点 target ，和一个整数值 k 。
//
// 返回到目标结点 target 距离为 k 的所有结点的值的列表。 答案可以以 任何顺序 返回。
//
//
//
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [3,5,1,6,2,0,8,null,null,7,4], target = 5, k = 2
// 输出：[7,4,1]
// 解释：所求结点为与目标结点（值为 5）距离为 2 的结点，值分别为 7，4，以及 1
//
//
// 示例 2:
//
//
// 输入: root = [1], target = 1, k = 3
// 输出: []
//
//
//
//
// 提示:
//
//
// 节点数在 [1, 500] 范围内
// 0 <= Node.val <= 500
// Node.val 中所有值 不同
// 目标结点 target 是树上的结点。
// 0 <= k <= 1000
//
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 618 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	visit := []int{}
	findDistanceK(root, target, k, &visit)
	return visit
}

func findDistanceK(root, target *TreeNode, K int, visit *[]int) int {
	if root == nil {
		return -1
	}
	if root == target {
		findChild(root, K, visit)
		return K - 1
	}

	leftDistance := findDistanceK(root.Left, target, K, visit)
	if leftDistance == 0 {
		// findChild(root, leftDistance, visit)
		*visit = append(*visit, root.Val)
	}
	if leftDistance > 0 {
		findChild(root.Right, leftDistance-1, visit)
		return leftDistance - 1
	}

	rightDistance := findDistanceK(root.Right, target, K, visit)
	if rightDistance == 0 {
		// findChild(root, rightDistance, visit)
		*visit = append(*visit, root.Val)
	}
	if rightDistance > 0 {
		findChild(root.Left, rightDistance-1, visit)
		return rightDistance - 1
	}
	return -1
}

func findChild(root *TreeNode, K int, visit *[]int) {
	if root == nil {
		return
	}
	if K == 0 {
		*visit = append(*visit, root.Val)
	} else {
		findChild(root.Left, K-1, visit)
		findChild(root.Right, K-1, visit)
	}
}

// leetcode submit region end(Prohibit modification and deletion)
