// 给你一个二叉树的根结点 root ，请返回出现次数最多的子树元素和。如果有多个元素出现的次数相同，返回所有出现次数最多的子树元素和（不限顺序）。
//
// 一个结点的 「子树元素和」 定义为以该结点为根的二叉树上所有结点的元素之和（包括结点本身）。
//
//
//
// 示例 1：
//
//
//
//
// 输入: root = [5,2,-3]
// 输出: [2,-3,4]
//
//
// 示例 2：
//
//
//
//
// 输入: root = [5,2,-5]
// 输出: [2]
//
//
//
//
// 提示:
//
//
// 节点数在 [1, 10⁴] 范围内
// -10⁵ <= Node.val <= 10⁵
//
//
// Related Topics 树 深度优先搜索 哈希表 二叉树 👍 220 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 同样是遍历每个节点作为根与子树元素的和，用一个map记录下来次数
// 返回出现子树和次数最多的和
func findFrequentTreeSum(root *TreeNode) []int {
	var res []int
	record := make(map[int]int, 0)
	dfs(root, record)
	max := 0
	for key, val := range record {
		if max == val {
			res = append(res, key)
		} else if max < val {
			max = val
			res = []int{key}
		}
	}
	return res
}
func dfs(root *TreeNode, record map[int]int) int {
	if root == nil {
		return 0
	}
	sum := root.Val + dfs(root.Left, record) + dfs(root.Right, record)
	if v, ok := record[sum]; ok {
		record[sum] = v + 1
	} else {
		record[sum] = 1
	}
	return sum
}

// leetcode submit region end(Prohibit modification and deletion)
