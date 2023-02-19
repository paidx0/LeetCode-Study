//给你二叉树的根节点 root ，返回其节点值的 锯齿形层序遍历 。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。
//
//
//
// 示例 1：
//
//
//输入：root = [3,9,20,null,null,15,7]
//输出：[[3],[20,9],[15,7]]
//
//
// 示例 2：
//
//
//输入：root = [1]
//输出：[[1]]
//
//
// 示例 3：
//
//
//输入：root = []
//输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [0, 2000] 内
// -100 <= Node.val <= 100
//
//
// Related Topics 树 广度优先搜索 二叉树 👍 736 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func zigzagLevelOrder(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}
	res := make([][]int, 0)
	level := []*TreeNode{pRoot}
	for i := 0; len(level) > 0; i++ {
		levelLen := len(level)
		res = append(res, make([]int, levelLen))
		temp := make([]*TreeNode, 0)
		for j := 0; j < levelLen; j++ {
			cur := level[j]
			if i%2 == 0 {
				res[i][j] = cur.Val
			} else {
				res[i][levelLen-1-j] = cur.Val
			}

			if cur.Left != nil {
				temp = append(temp, cur.Left)
			}
			if cur.Right != nil {
				temp = append(temp, cur.Right)
			}

		}
		level = temp
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
