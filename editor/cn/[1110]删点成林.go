// 给出二叉树的根节点 root，树上每个节点都有一个不同的值。
//
// 如果节点值在 to_delete 中出现，我们就把该节点从树上删去，最后得到一个森林（一些不相交的树构成的集合）。
//
// 返回森林中的每棵树。你可以按任意顺序组织答案。
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [1,2,3,4,5,6,7], to_delete = [3,5]
// 输出：[[1,2,null,4],[6],[7]]
//
//
// 示例 2：
//
//
// 输入：root = [1,2,4,null,3], to_delete = [3]
// 输出：[[1,2,4]]
//
//
//
//
// 提示：
//
//
// 树中的节点数最大为 1000。
// 每个节点都有一个介于 1 到 1000 之间的值，且各不相同。
// to_delete.length <= 1000
// to_delete 包含一些从 1 到 1000、各不相同的值。
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 203 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	res, deleteMap := []*TreeNode{}, map[int]bool{}
	for _, v := range to_delete {
		deleteMap[v] = true
	}
	if dfs(root, &res, deleteMap) != nil {
		res = append(res, root)
	}
	return res
}
func dfs(root *TreeNode, res *[]*TreeNode, deleteMap map[int]bool) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = dfs(root.Left, res, deleteMap)
	root.Right = dfs(root.Right, res, deleteMap)
	if deleteMap[root.Val] {
		if root.Left != nil {
			*res = append(*res, root.Left)
		}
		if root.Right != nil {
			*res = append(*res, root.Right)
		}
		return nil
	}
	return root
}

// leetcode submit region end(Prohibit modification and deletion)
