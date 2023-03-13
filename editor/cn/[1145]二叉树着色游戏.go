// 有两位极客玩家参与了一场「二叉树着色」的游戏。游戏中，给出二叉树的根节点 root，树上总共有 n 个节点，且 n 为奇数，其中每个节点上的值从 1 到
// n 各不相同。
//
// 最开始时：
//
//
// 「一号」玩家从 [1, n] 中取一个值 x（1 <= x <= n）；
// 「二号」玩家也从 [1, n] 中取一个值 y（1 <= y <= n）且 y != x。
//
//
// 「一号」玩家给值为 x 的节点染上红色，而「二号」玩家给值为 y 的节点染上蓝色。
//
// 之后两位玩家轮流进行操作，「一号」玩家先手。每一回合，玩家选择一个被他染过色的节点，将所选节点一个 未着色 的邻节点（即左右子节点、或父节点）进行染色（「
// 一号」玩家染红色，「二号」玩家染蓝色）。
//
// 如果（且仅在此种情况下）当前玩家无法找到这样的节点来染色时，其回合就会被跳过。
//
// 若两个玩家都没有可以染色的节点时，游戏结束。着色节点最多的那位玩家获得胜利 ✌️。
//
// 现在，假设你是「二号」玩家，根据所给出的输入，假如存在一个 y 值可以确保你赢得这场游戏，则返回 true ；若无法获胜，就请返回 false 。
//
// 示例 1 ：
//
//
// 输入：root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
// 输出：true
// 解释：第二个玩家可以选择值为 2 的节点。
//
// 示例 2 ：
//
//
// 输入：root = [1,2,3], n = 3, x = 1
// 输出：false
//
//
//
//
// 提示：
//
//
// 树中节点数目为 n
// 1 <= x <= n <= 100
// n 是奇数
// 1 <= Node.val <= n
// 树中所有值 互不相同
//
//
// Related Topics 树 深度优先搜索 二叉树 👍 200 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 红色把树分成三部分，左子树，右子树，和父节点那边
// 要想赢，那就只要蓝色堵在红色的父节点处，那红色只能是左右两部分
// 只需要大于这两部分就能赢
// 还有种情况是红色就在根节点，分成两部分，那蓝色只需要能选择更多的一边就可以堵住红色
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var left, right int
	dfs(root, n, x, &left, &right)
	blue := n - left - right - 1
	return blue > n/2 || left > n/2 || right > n/2
}
func dfs(root *TreeNode, n int, x int, left, right *int) int {
	if root == nil {
		return 0
	}
	l := dfs(root.Left, n, x, left, right)
	r := dfs(root.Right, n, x, left, right)
	if root.Val == x {
		// 找到左右部分
		*left = l
		*right = r
	}
	return l + r + 1
}

// leetcode submit region end(Prohibit modification and deletion)
