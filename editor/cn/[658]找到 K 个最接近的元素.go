// 给定一个 排序好 的数组 arr ，两个整数 k 和 x ，从数组中找到最靠近 x（两数之差最小）的 k 个数。返回的结果必须要是按升序排好的。
//
// 整数 a 比整数 b 更接近 x 需要满足：
//
//
// |a - x| < |b - x| 或者
// |a - x| == |b - x| 且 a < b
//
//
//
//
// 示例 1：
//
//
// 输入：arr = [1,2,3,4,5], k = 4, x = 3
// 输出：[1,2,3,4]
//
//
// 示例 2：
//
//
// 输入：arr = [1,2,3,4,5], k = 4, x = -1
// 输出：[1,2,3,4]
//
//
//
//
// 提示：
//
//
// 1 <= k <= arr.length
// 1 <= arr.length <= 10⁴
//
// arr 按 升序 排列
// -10⁴ <= arr[i], x <= 10⁴
//
//
// Related Topics 数组 双指针 二分查找 排序 滑动窗口 堆（优先队列） 👍 480 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// 二分查找找左边边界
// 左边边界最长的搜索范围[0,len(arr)-k]，也就是在最右边时
func findClosestElements(arr []int, k int, x int) []int {
	low, high := 0, len(arr)-k
	for low < high {
		mid := low + (high-low)>>1
		if x-arr[mid] > arr[mid+k]-x {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return arr[low : low+k]
}

// leetcode submit region end(Prohibit modification and deletion)
