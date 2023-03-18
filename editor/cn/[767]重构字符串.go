// 给定一个字符串 s ，检查是否能重新排布其中的字母，使得两相邻的字符不同。
//
// 返回 s 的任意可能的重新排列。若不可行，返回空字符串 "" 。
//
//
//
// 示例 1:
//
//
// 输入: s = "aab"
// 输出: "aba"
//
//
// 示例 2:
//
//
// 输入: s = "aaab"
// 输出: ""
//
//
//
//
// 提示:
//
//
// 1 <= s.length <= 500
// s 只包含小写字母
//
//
// Related Topics 贪心 哈希表 字符串 计数 排序 堆（优先队列） 👍 466 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// 可以用堆，每次从最大堆中取出两个最大的拼接
// 也可以用451那题，先把字母长度按照从大到小排好序
// 然后双指针从头和中间开始取两个拼接，如果能走到最后一个那就正常返回
import (
	"sort"
)

func reorganizeString(S string) string {
	fs := frequencySort451(S)
	if fs == "" {
		return ""
	}
	bs := []byte(fs)
	ans := ""
	j := (len(bs)-1)/2 + 1
	// 双指针
	for i := 0; i <= (len(bs)-1)/2; i++ {
		ans += string(bs[i])
		if j < len(bs) {
			ans += string(bs[j])
		}
		j++
	}
	return ans
}
func frequencySort451(s string) string {
	sMap := map[byte]int{}
	cMap := map[int][]byte{}
	sb := []byte(s)
	// 统计长度，超过一半直接返回
	for _, b := range sb {
		sMap[b]++
		if sMap[b] > (len(sb)+1)/2 {
			return ""
		}
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}
	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]byte, 0)
	// 拼接字符串返回
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			res = append(res, bytes.Repeat([]byte{cMap[k][i]}, k)...)
		}
	}

	return string(res)
}

// leetcode submit region end(Prohibit modification and deletion)
