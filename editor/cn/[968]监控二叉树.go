// 给定一个二叉树，我们在树的节点上安装摄像头。
//
// 节点上的每个摄影头都可以监视其父对象、自身及其直接子对象。
//
// 计算监控树的所有节点所需的最小摄像头数量。
//
//
//
// 示例 1：
//
//
//
// 输入：[0,0,null,0,0]
// 输出：1
// 解释：如图所示，一台摄像头足以监控所有节点。
//
//
// 示例 2：
//
//
//
// 输入：[0,0,null,0,null,0,null,null,0]
// 输出：2
// 解释：需要至少两个摄像头来监视树的所有节点。 上图显示了摄像头放置的有效位置之一。
//
//
// 提示：
//
//
// 给定树的节点数的范围是 [1, 1000]。
// 每个节点的值都是 0。
//
//
// Related Topics 树 深度优先搜索 动态规划 二叉树 👍 561 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

const (
	isLeaf       int = iota // 叶子
	parentofLeaf            // 叶子的父
	isMonitored             // 被覆盖
)

// 后序遍历，首先所有的叶子节点的父肯定是可以给一个的，可能可以覆盖到4个点
// 那么叶子的父的父自然是可以被覆盖到的，他的父开始又是一次循环,就这样往上层移动，记录状态
func minCameraCover(root *TreeNode) int {
	var res int
	if dfs(root, &res) == isLeaf {
		// 这里还要另外检查一下返回来的父节点，可能没有被覆盖
		res++
	}
	return res
}
func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 2
	}
	left := dfs(root.Left, res)
	right := dfs(root.Right, res)
	switch {
	case left == isLeaf || right == isLeaf:
		// 子节点的父，必然要加监控
		*res++
		return parentofLeaf
	case left == parentofLeaf || right == parentofLeaf:
		// 子节点的父的父，被覆盖
		return isMonitored
	default:
		// left == right == nil，那就是叶子
		return isLeaf
	}
}

// leetcode submit region end(Prohibit modification and deletion)
