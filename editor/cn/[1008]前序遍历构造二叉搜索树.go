//给定一个整数数组，它表示BST(即 二叉搜索树 )的 先序遍历 ，构造树并返回其根。
//
// 保证 对于给定的测试用例，总是有可能找到具有给定需求的二叉搜索树。
//
// 二叉搜索树 是一棵二叉树，其中每个节点， Node.left 的任何后代的值 严格小于 Node.val , Node.right 的任何后代的值 严格大
//于 Node.val。
//
// 二叉树的 前序遍历 首先显示节点的值，然后遍历Node.left，最后遍历Node.right。
//
//
//
// 示例 1：
//
//
//
//
//输入：preorder = [8,5,1,7,10,12]
//输出：[8,5,10,1,7,null,12]
//
//
// 示例 2:
//
//
//输入: preorder = [1,3]
//输出: [1,null,3]
//
//
//
//
// 提示：
//
//
// 1 <= preorder.length <= 100
// 1 <= preorder[i] <= 10^8
// preorder 中的值 互不相同
//
//
//
//
// Related Topics 栈 树 二叉搜索树 数组 二叉树 单调栈 👍 246 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
	"sort"
)

// 根据前序和中序得到原树
func bstFromPreorder(preorder []int) *TreeNode {
	inorder := make([]int, len(preorder))
	copy(inorder, preorder)
	sort.Ints(inorder)
	return buildTree(preorder, inorder)
}

// 方法一
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	inPos := make(map[int]int)
//	for i := 0; i < len(inorder); i++ {
//		inPos[inorder[i]] = i
//	}
//	return buildPreIn2TreeDFS(preorder, 0, len(preorder)-1, 0, inPos)
//}
//func buildPreIn2TreeDFS(pre []int, preStart int, preEnd int, inStart int, inPos map[int]int) *TreeNode {
//	if preStart > preEnd {
//		return nil
//	}
//	root := &TreeNode{Val: pre[preStart]}
//	rootIdx := inPos[pre[preStart]]
//	leftLen := rootIdx - inStart
//	root.Left = buildPreIn2TreeDFS(pre, preStart+1, preStart+leftLen, inStart, inPos)
//	root.Right = buildPreIn2TreeDFS(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
//	return root
//}

// 方法二
func buildTree(preorder []int, inorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildPreIn2TreeDFS(preorder, 0, len(preorder)-1, 0, inPos)
}
func buildPreIn2TreeDFS(pre []int, preStart int, preEnd int, inStart int, inPos map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	root := &TreeNode{Val: pre[preStart]}
	rootIdx := inPos[pre[preStart]]
	leftLen := rootIdx - inStart
	root.Left = buildPreIn2TreeDFS(pre, preStart+1, preStart+leftLen, inStart, inPos)
	root.Right = buildPreIn2TreeDFS(pre, preStart+leftLen+1, preEnd, rootIdx+1, inPos)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
