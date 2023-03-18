// 给定两个以 升序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
//
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
//
// 请找到和最小的 k 个数对 (u1,v1), (u2,v2) ... (uk,vk) 。
//
//
//
// 示例 1:
//
//
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [1,2],[1,4],[1,6]
// 解释: 返回序列中的前 3 对数：
//     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
//
//
// 示例 2:
//
//
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [1,1],[1,1]
// 解释: 返回序列中的前 2 对数：
//      [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
//
//
// 示例 3:
//
//
// 输入: nums1 = [1,2], nums2 = [3], k = 3
// 输出: [1,3],[2,3]
// 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]
//
//
//
//
// 提示:
//
//
// 1 <= nums1.length, nums2.length <= 10⁵
// -10⁹ <= nums1[i], nums2[i] <= 10⁹
// nums1 和 nums2 均为升序排列
// 1 <= k <= 1000
//
//
// Related Topics 数组 堆（优先队列） 👍 466 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// import (
// 	"sort"
// )
// 暴力一遍笛卡尔积
// func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
// 	size1, size2, res := len(nums1), len(nums2), [][]int{}
// 	for i := 0; i < size1; i++ {
// 		for j := 0; j < size2; j++ {
// 			res = append(res, []int{nums1[i], nums2[j]})
// 		}
// 	}
// 	sort.Slice(res, func(i, j int) bool {
// 		return res[i][0]+res[i][1] < res[j][0]+res[j][1]
// 	})
// 	if len(nums1)*len(nums2) < k {
// 		k = len(nums1) * len(nums2)
// 	}
// 	if len(res) >= k {
// 		return res[:k]
// 	}
// 	return res
// }

// import (
// 	"container/heap"
// )
// 最小堆
// func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
// 	result, h := [][]int{}, &minHeap{}
//
// 	if len(nums1)*len(nums2) < k {
// 		k = len(nums1) * len(nums2)
// 	}
//
// 	for _, num := range nums1 {
// 		heap.Push(h, []int{num, nums2[0], 0})
// 	}
//
// 	for len(result) < k {
// 		min := heap.Pop(h).([]int)
// 		result = append(result, min[:2])
// 		if min[2] < len(nums2)-1 {
// 			heap.Push(h, []int{min[0], nums2[min[2]+1], min[2] + 1})
// 		}
// 	}
// 	return result
// }
//
// type minHeap [][]int
//
// func (h minHeap) Len() int            { return len(h) }
// func (h minHeap) Less(i, j int) bool  { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
// func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
// func (h *minHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
// func (h *minHeap) Pop() interface{} {old := *h; n := len(old); x := old[n-1]; *h = old[0: n-1]; return x}

import (
	"container/heap"
)

func kSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := hp{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}

type pair struct{ i, j int }
type hp struct {
	data         []pair
	nums1, nums2 []int
}

func (h hp) Len() int { return len(h.data) }
func (h hp) Less(i, j int) bool {
	a, b := h.data[i], h.data[j]
	return h.nums1[a.i]+h.nums2[a.j] < h.nums1[b.i]+h.nums2[b.j]
}
func (h hp) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *hp) Push(v interface{}) { h.data = append(h.data, v.(pair)) }
func (h *hp) Pop() interface{}   { a := h.data; v := a[len(a)-1]; h.data = a[:len(a)-1]; return v }

// leetcode submit region end(Prohibit modification and deletion)
