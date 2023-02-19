//给你 root1 和 root2 这两棵二叉搜索树。请你返回一个列表，其中包含 两棵树 中的所有整数并按 升序 排序。.
//
//
//
// 示例 1：
//
//
//
//
//输入：root1 = [2,1,4], root2 = [1,0,3]
//输出：[0,1,1,2,3,4]
//
//
// 示例 2：
//
//
//
//
//输入：root1 = [1,null,8], root2 = [8,1]
//输出：[1,1,8,8]
//
//
//
//
// 提示：
//
//
// 每棵树的节点数在 [0, 5000] 范围内
// -10⁵ <= Node.val <= 10⁵
//
//
// Related Topics 树 深度优先搜索 二叉搜索树 二叉树 排序 👍 163 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 中序遍历拿到两个有序数组，后面就是数组的合并了，参考88题
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var list1, list2, res []int
	dfs(root1, &list1)
	dfs(root2, &list2)
	m, n := len(list1), len(list2)
	// 扩容再合并
	list1 = append(list1, make([]int, n)...)
	res = merge(list1, m, list2, n)
	return res
}
func dfs(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, list)
	*list = append(*list, root.Val)
	dfs(root.Right, list)
}
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	for p := m + n; m > 0 && n > 0; p-- {
		if nums1[m-1] <= nums2[n-1] {
			nums1[p-1] = nums2[n-1]
			n--
		} else {
			nums1[p-1] = nums1[m-1]
			m--
		}
	}
	for ; n > 0; n-- {
		nums1[n-1] = nums2[n-1]
	}
	return nums1
}

//leetcode submit region end(Prohibit modification and deletion)
