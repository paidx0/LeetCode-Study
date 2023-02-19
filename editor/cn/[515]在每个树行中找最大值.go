//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。
//
//
//
// 示例1：
//
//
//
//
//输入: root = [1,3,2,5,3,null,9]
//输出: [1,3,9]
//
//
// 示例2：
//
//
//输入: root = [1,2,3]
//输出: [1,3]
//
//
//
//
// 提示：
//
//
// 二叉树的节点个数的范围是 [0,10⁴]
//
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 二叉树 👍 298 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree cur.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	level := []*TreeNode{root}
	var res []int
	for len(level) > 0 {
		levellen := len(level)
		max := math.MinInt32
		for i := 0; i < levellen; i++ {
			cur := level[0]
			level = level[1:]
			if cur.Val > max {
				max = cur.Val
			}
			if cur.Left != nil {
				level = append(level, cur.Left)
			}
			if cur.Right != nil {
				level = append(level, cur.Right)
			}
		}
		res = append(res, max)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
