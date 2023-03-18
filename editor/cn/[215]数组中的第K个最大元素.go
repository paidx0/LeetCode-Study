// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
//
// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
//
//
//
// 示例 1:
//
//
// 输入: [3,2,1,5,6,4], k = 2
// 输出: 5
//
//
// 示例 2:
//
//
// 输入: [3,2,3,1,2,4,5,5,6], k = 4
// 输出: 4
//
//
//
// 提示：
//
//
// 1 <= k <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
// Related Topics 数组 分治 快速选择 排序 堆（优先队列） 👍 2104 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// import (
// 	"sort"
// )
// func findKthLargest(nums []int, k int) int {
// 	sort.Ints(nums)
// 	return nums[len(nums)-k]
// }

// 堆排序
func findKthLargest(nums []int, k int) int {
	arrLen := len(nums)
	// 建堆
	for i := arrLen / 2; i >= 0; i-- {
		heapify(nums, i, arrLen)
	}
	// 维护
	for i := arrLen - 1; i >= 0; i-- {
		swap(nums, 0, i)
		arrLen--
		heapify(nums, 0, arrLen)
	}
	return nums[k-1]
}
func heapify(arr []int, i, arrLen int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < arrLen && arr[left] < arr[largest] {
		largest = left
	}
	if right < arrLen && arr[right] < arr[largest] {
		largest = right
	}
	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, arrLen)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// leetcode submit region end(Prohibit modification and deletion)
