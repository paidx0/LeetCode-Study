// 在一个仓库里，有一排条形码，其中第 i 个条形码为 barcodes[i]。
//
// 请你重新排列这些条形码，使其中任意两个相邻的条形码不能相等。 你可以返回任何满足该要求的答案，此题保证存在答案。
//
//
//
// 示例 1：
//
//
// 输入：barcodes = [1,1,1,2,2,2]
// 输出：[2,1,2,1,2,1]
//
//
// 示例 2：
//
//
// 输入：barcodes = [1,1,1,1,2,2,3,3]
// 输出：[1,3,1,3,2,1,2,1]
//
//
//
// 提示：
//
//
// 1 <= barcodes.length <= 10000
// 1 <= barcodes[i] <= 10000
//
//
// Related Topics 贪心 数组 哈希表 计数 排序 堆（优先队列） 👍 98 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// 和767一样的，排序后从开头和中间开始取拼接
import (
	"sort"
)

func rearrangeBarcodes(barcodes []int) []int {
	bfs := FrequencySort451(barcodes)
	if len(bfs) == 0 {
		return []int{}
	}
	res := []int{}
	j := (len(bfs)-1)/2 + 1
	for i := 0; i <= (len(bfs)-1)/2; i++ {
		res = append(res, bfs[i])
		if j < len(bfs) {
			res = append(res, bfs[j])
		}
		j++
	}
	return res
}
func FrequencySort451(s []int) []int {
	sMap := map[int]int{}
	cMap := map[int][]int{}
	for _, b := range s {
		sMap[b]++
	}
	for key, value := range sMap {
		cMap[value] = append(cMap[value], key)
	}
	var keys []int
	for k := range cMap {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	res := make([]int, 0)
	for _, k := range keys {
		for i := 0; i < len(cMap[k]); i++ {
			for j := 0; j < k; j++ {
				res = append(res, cMap[k][i])
			}
		}
	}
	return res
}

// leetcode submit region end(Prohibit modification and deletion)
