// 给你一个按 非递减顺序 排列的整数数组 nums 。
//
// 请你判断是否能在将 nums 分割成 一个或多个子序列 的同时满足下述 两个 条件：
//
//
//
//
// 每个子序列都是一个 连续递增序列（即，每个整数 恰好 比前一个整数大 1 ）。
// 所有子序列的长度 至少 为 3 。
//
//
//
//
// 如果可以分割 nums 并满足上述条件，则返回 true ；否则，返回 false 。
//
//
//
// 示例 1：
//
//
// 输入：nums = [1,2,3,3,4,5]
// 输出：true
// 解释：nums 可以分割成以下子序列：
// [1,2,3,3,4,5] --> 1, 2, 3
// [1,2,3,3,4,5] --> 3, 4, 5
//
//
// 示例 2：
//
//
// 输入：nums = [1,2,3,3,4,4,5,5]
// 输出：true
// 解释：nums 可以分割成以下子序列：
// [1,2,3,3,4,4,5,5] --> 1, 2, 3, 4, 5
// [1,2,3,3,4,4,5,5] --> 3, 4, 5
//
//
// 示例 3：
//
//
// 输入：nums = [1,2,3,4,4,5]
// 输出：false
// 解释：无法将 nums 分割成长度至少为 3 的连续递增子序列。
//
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁴
// -1000 <= nums[i] <= 1000
// nums 按非递减顺序排列
//
//
// Related Topics 贪心 数组 哈希表 堆（优先队列） 👍 420 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

import (
	"container/heap"
)

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }

func isPossible(nums []int) bool {
	lens := map[int]*hp{}
	// 每个位置维护一个队列数组，存数组长度
	for _, v := range nums {
		if lens[v] == nil {
			lens[v] = new(hp)
		}
		// 前一个数字存在就把当前数字也加进去
		// 不存在则新开一个队列数组，从长度1开始
		if h := lens[v-1]; h != nil {
			prevLen := h.pop()
			// 没有重复数字，则删除前一个队列，加到后面队列去
			if h.Len() == 0 {
				delete(lens, v-1)
			}
			lens[v].push(prevLen + 1)
		} else {
			lens[v].push(1)
		}
	}
	for _, h := range lens {
		if h.IntSlice[0] < 3 {
			return false
		}
	}
	return true
}

// 贪心
// func isPossible(nums []int) bool {
// 	left := map[int]int{}
// 	for _, v := range nums {
// 		left[v]++
// 	}
// 	endCnt := map[int]int{}
// 	for _, v := range nums {
// 		if left[v] == 0 {
// 			continue
// 		}
// 		if endCnt[v-1] > 0 { // 若存在以 v-1 结尾的连续子序列，则将 v 加到其末尾
// 			left[v]--
// 			endCnt[v-1]--
// 			endCnt[v]++
// 		} else if left[v+1] > 0 && left[v+2] > 0 { // 否则，生成一个长度为 3 的连续子序列
// 			left[v]--
// 			left[v+1]--
// 			left[v+2]--
// 			endCnt[v+2]++
// 		} else { // 无法生成
// 			return false
// 		}
// 	}
// 	return true
// }

// leetcode submit region end(Prohibit modification and deletion)
