//给定一个非空特殊的二叉树，每个节点都是正数，并且每个节点的子节点数量只能为 2 或 0。如果一个节点有两个子节点的话，那么该节点的值等于两个子节点中较小的一
//个。
//
// 更正式地说，即 root.val = min(root.left.val, root.right.val) 总成立。
//
// 给出这样的一个二叉树，你需要输出所有节点中的 第二小的值 。
//
// 如果第二小的值不存在的话，输出 -1 。
//
//
//
// 示例 1：
//
//
//输入：root = [2,2,5,null,null,5,7]
//输出：5
//解释：最小的值是 2 ，第二小的值是 5 。
//
//
// 示例 2：
//
//
//输入：root = [2,2,2]
//输出：-1
//解释：最小的值是 2, 但是不存在第二小的值。
//
//
//
//
// 提示：
//
//
// 树中节点数目在范围 [1, 25] 内
// 1 <= Node.val <= 2³¹ - 1
// 对于树中每个节点 root.val == min(root.left.val, root.right.val)
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 285 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 层序遍历，只需要找下面层中不等于根节点然后最小的就行
// 只有两种数字也是 -1
func findSecondMinimumValue(root *TreeNode) int {
	ans := -1
	rootVal := root.Val
	dfs(root, &ans, rootVal)
	return ans
}
func dfs(node *TreeNode, ans *int, rootVal int) {
	if node == nil || *ans != -1 && node.Val >= *ans {
		return
	}
	if node.Val > rootVal {
		*ans = node.Val
	}
	dfs(node.Left, ans, rootVal)
	dfs(node.Right, ans, rootVal)
}

//leetcode submit region end(Prohibit modification and deletion)
