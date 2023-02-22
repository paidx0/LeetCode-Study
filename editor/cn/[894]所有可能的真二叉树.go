// 给你一个整数 n ，请你找出所有可能含 n 个节点的 真二叉树 ，并以列表形式返回。答案中每棵树的每个节点都必须符合 Node.val == 0 。
//
// 答案的每个元素都是一棵真二叉树的根节点。你可以按 任意顺序 返回最终的真二叉树列表。
//
// 真二叉树 是一类二叉树，树中每个节点恰好有 0 或 2 个子节点。
//
//
//
// 示例 1：
//
//
// 输入：n = 7
// 输出：[[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0
// ,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
//
//
// 示例 2：
//
//
// 输入：n = 3
// 输出：[[0,0,0]]
//
//
//
//
// 提示：
//
//
// 1 <= n <= 20
//
//
// Related Topics 树 递归 记忆化搜索 动态规划 二叉树 👍 286 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 满二叉树一定是奇数，偶数直接返回nil
// 遍历存下每种节点数的情况
// 左边有 j个时，右边有 i-1-j个
func allPossibleFBT(n int) []*TreeNode {
	if n%2 == 0 {
		return nil
	}
	nodeMap := make(map[int][]*TreeNode, n)
	for i := 1; i <= n; i = i + 2 {
		if i == 1 {
			nodeMap[1] = append(nodeMap[1], &TreeNode{Val: 0})
			continue
		}
		for j := 1; j < i; j = j + 2 {
			for _, m := range nodeMap[j] {
				for _, n := range nodeMap[i-1-j] {
					nowNode := &TreeNode{
						Val:   0,
						Left:  &TreeNode{Val: 0, Left: m.Left, Right: m.Right},
						Right: &TreeNode{Val: 0, Left: n.Left, Right: n.Right},
					}
					nodeMap[i] = append(nodeMap[i], nowNode)
				}
			}
		}
	}
	return nodeMap[n]
}

// leetcode submit region end(Prohibit modification and deletion)
