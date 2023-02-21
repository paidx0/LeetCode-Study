//给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。可以按 任意顺序 返回答案。
//
//
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：[[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
//
//
//
//
// 示例 2：
//
//
//输入：n = 1
//输出：[[1]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 8
//
//
// Related Topics 树 二叉搜索树 动态规划 回溯 二叉树 👍 1384 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return generateBSTree(1, n)
}
func generateBSTree(start, end int) []*TreeNode {
	tree := []*TreeNode{}
	if start > end {
		tree = append(tree, nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := generateBSTree(start, i-1)
		right := generateBSTree(i+1, end)
		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: i, Left: l, Right: r}
				tree = append(tree, root)
			}
		}
	}
	return tree
}

//leetcode submit region end(Prohibit modification and deletion)
