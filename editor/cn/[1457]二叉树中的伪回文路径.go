// 给你一棵二叉树，每个节点的值为 1 到 9 。我们称二叉树中的一条路径是 「伪回文」的，当它满足：路径经过的所有节点值的排列中，存在一个回文序列。
//
// 请你返回从根到叶子节点的所有路径中 伪回文 路径的数目。
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [2,3,1,3,1,null,1]
// 输出：2
// 解释：上图为给定的二叉树。总共有 3 条从根到叶子的路径：红色路径 [2,3,3] ，绿色路径 [2,1,1] 和路径 [2,3,1] 。
//     在这些路径中，只有红色和绿色的路径是伪回文路径，因为红色路径 [2,3,3] 存在回文排列 [3,2,3] ，绿色路径 [2,1,1] 存在回文排
// 列 [1,2,1] 。
//
//
// 示例 2：
//
//
//
//
// 输入：root = [2,1,1,1,3,null,null,null,null,null,1]
// 输出：1
// 解释：上图为给定二叉树。总共有 3 条从根到叶子的路径：绿色路径 [2,1,1] ，路径 [2,1,3,1] 和路径 [2,1] 。
//     这些路径中只有绿色路径是伪回文路径，因为 [2,1,1] 存在回文排列 [1,2,1] 。
//
//
// 示例 3：
//
//
// 输入：root = [9]
// 输出：1
//
//
//
//
// 提示：
//
//
// 给定二叉树的节点数目在范围 [1, 10⁵] 内
// 1 <= Node.val <= 9
//
//
// Related Topics 位运算 树 深度优先搜索 广度优先搜索 二叉树 👍 60 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 位运算，要么成对异或后能消除等于0，要么就是只剩下一位为1，减1做&验证
func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	dfs(root, &res, 0)
	return res
}
func dfs(root *TreeNode, res *int, temp int) {
	n := temp ^ (1 << root.Val)
	if root.Left == nil && root.Right == nil {
		if n == 0 || (n&(n-1)) == 0 {
			*res += 1
		}
		return
	}
	if root.Left != nil {
		dfs(root.Left, res, n)
	}
	if root.Right != nil {
		dfs(root.Right, res, n)
	}
}

// leetcode submit region end(Prohibit modification and deletion)
