// 序列化是将数据结构或对象转换为一系列位的过程，以便它可以存储在文件或内存缓冲区中，或通过网络连接链路传输，以便稍后在同一个或另一个计算机环境中重建。
//
// 设计一个算法来序列化和反序列化 二叉搜索树 。 对序列化/反序列化算法的工作方式没有限制。 您只需确保二叉搜索树可以序列化为字符串，并且可以将该字符串反序
// 列化为最初的二叉搜索树。
//
// 编码的字符串应尽可能紧凑。
//
//
//
// 示例 1：
//
//
// 输入：root = [2,1,3]
// 输出：[2,1,3]
//
//
// 示例 2：
//
//
// 输入：root = []
// 输出：[]
//
//
//
//
// 提示：
//
//
// 树中节点数范围是 [0, 10⁴]
// 0 <= Node.val <= 10⁴
// 题目数据 保证 输入的树是一棵二叉搜索树。
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 设计 二叉搜索树 字符串 二叉树 👍 407 👎 0

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
	"math"
	"strconv"
	"strings"
)

type Codec struct {
	idx int
}

func Constructor() Codec {
	return Codec{idx: 0}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var preOrder func(node *TreeNode) string
	preOrder = func(node *TreeNode) string {
		if node == nil {
			return ""
		}
		return strconv.Itoa(node.Val) + "," + preOrder(node.Left) + preOrder(node.Right)
	}
	return preOrder(root)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	var dfs func(arr []string, minimum int, maximum int) *TreeNode
	dfs = func(arr []string, minimum int, maximum int) *TreeNode {
		if this.idx >= len(arr) || len(arr[this.idx]) == 0 {
			return nil
		}
		curr, _ := strconv.Atoi(arr[this.idx])
		if curr < minimum || curr > maximum {
			return nil
		}
		this.idx++
		node := &TreeNode{Val: curr}
		node.Left = dfs(arr, minimum, curr)
		node.Right = dfs(arr, curr, maximum)
		return node
	}
	return dfs(strings.Split(data, ","), math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// leetcode submit region end(Prohibit modification and deletion)
