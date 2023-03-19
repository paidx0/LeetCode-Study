// 给你一个字符数组 letters，该数组按非递减顺序排序，以及一个字符 target。letters 里至少有两个不同的字符。
//
// 返回 letters 中大于 target 的最小的字符。如果不存在这样的字符，则返回 letters 的第一个字符。
//
//
//
// 示例 1：
//
//
// 输入: letters = ["c", "f", "j"]，target = "a"
// 输出: "c"
// 解释：letters 中字典上比 'a' 大的最小字符是 'c'。
//
// 示例 2:
//
//
// 输入: letters = ["c","f","j"], target = "c"
// 输出: "f"
// 解释：letters 中字典顺序上大于 'c' 的最小字符是 'f'。
//
// 示例 3:
//
//
// 输入: letters = ["x","x","y","y"], target = "z"
// 输出: "x"
// 解释：letters 中没有一个字符在字典上大于 'z'，所以我们返回 letters[0]。
//
//
//
// 提示：
//
//
// 2 <= letters.length <= 10⁴
// letters[i] 是一个小写字母
// letters 按非递减顺序排序
// letters 最少包含两个不同的字母
// target 是一个小写字母
//
//
// Related Topics 数组 二分查找 👍 256 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
func nextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)-1
	for low <= high {
		mid := low + (high-low)>>1
		if letters[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	// 可能在最后面找到target，所以这个时候是没有符合的，返回第一个值
	find := letters[low%len(letters)]
	if find <= target {
		return letters[0]
	}
	return find
}

// leetcode submit region end(Prohibit modification and deletion)
