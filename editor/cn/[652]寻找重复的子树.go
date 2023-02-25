// 给你一棵二叉树的根节点 root ，返回所有 重复的子树 。
//
// 对于同一类的重复子树，你只需要返回其中任意 一棵 的根结点即可。
//
// 如果两棵树具有 相同的结构 和 相同的结点值 ，则认为二者是 重复 的。
//
//
//
// 示例 1：
//
//
//
//
// 输入：root = [1,2,3,4,null,2,4,null,null,4]
// 输出：[[2,4],[4]]
//
// 示例 2：
//
//
//
//
// 输入：root = [2,1,1]
// 输出：[[1]]
//
// 示例 3：
//
//
//
//
// 输入：root = [2,2,2,3,null,3,null]
// 输出：[[2,3],[3]]
//
//
//
// 提示：
//
//
// 树中的结点数在 [1, 5000] 范围内。
// -200 <= Node.val <= 200
//
//
// Related Topics 树 深度优先搜索 哈希表 二叉树 👍 664 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	"strconv"
)

var ump map[string]int

// 前序遍历，以字符串形式序列化，保存在一个map中，如果出现次数超过2次，就记录一次
// 因为是从下层开始添加序列化，所以最后需要翻转一次
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ump = make(map[string]int)
	var ans []*TreeNode
	dfs(root, &ans)
	return reverse(ans)
}
func dfs(root *TreeNode, ans *[]*TreeNode) string {
	if root == nil {
		return ""
	}
	ss := strconv.Itoa(root.Val) + " " + dfs(root.Left, ans) + " " + dfs(root.Right, ans)
	ump[ss]++
	if ump[ss] == 2 {
		*ans = append(*ans, root)
	}
	return ss
}
func reverse(ans []*TreeNode) []*TreeNode {
	n := len(ans)
	for i := 0; i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
