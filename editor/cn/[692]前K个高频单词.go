// 给定一个单词列表 words 和一个整数 k ，返回前 k 个出现次数最多的单词。
//
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率， 按字典顺序 排序。
//
//
//
// 示例 1：
//
//
// 输入: words = ["i", "love", "leetcode", "i", "love", "coding"], k = 2
// 输出: ["i", "love"]
// 解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
//
//
// 示例 2：
//
//
// 输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"],
// k = 4
// 输出: ["the", "is", "sunny", "day"]
// 解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
//
//
//
//
// 注意：
//
//
// 1 <= words.length <= 500
// 1 <= words[i] <= 10
// words[i] 由小写英文字母组成。
// k 的取值范围是 [1, 不同 words[i] 的数量]
//
//
//
//
// 进阶：尝试以 O(n log k) 时间复杂度和 O(n) 空间复杂度解决。
//
// Related Topics 字典树 哈希表 字符串 桶排序 计数 排序 堆（优先队列） 👍 529 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
import (
	"container/heap"
)

// 和703一样
// 堆里放最大的几个，维护一个最小堆，堆顶则是第k大
func topKFrequent(words []string, k int) []string {
	m := make(map[string]int, 0)
	for _, word := range words {
		m[word]++
	}
	pq := new(PQ)
	for w, c := range m {
		heap.Push(pq, &wordCount{w, c})
		if pq.Len() > k {
			heap.Pop(pq)
		}
	}
	// 从高到低返回
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		wc := heap.Pop(pq).(*wordCount)
		res[i] = wc.word
	}
	return res
}

type wordCount struct {
	word string
	cnt  int
}
type PQ []*wordCount

func (pq PQ) Len() int      { return len(pq) }
func (pq PQ) Swap(i, j int) { pq[i], pq[j] = pq[j], pq[i] }
func (pq PQ) Less(i, j int) bool {
	// 单词大的往前放，后面取出pop时是反的往后面放
	if pq[i].cnt == pq[j].cnt {
		return pq[i].word > pq[j].word
	}
	return pq[i].cnt < pq[j].cnt
}
func (pq *PQ) Push(x interface{}) { tmp := x.(*wordCount); *pq = append(*pq, tmp) }
func (pq *PQ) Pop() interface{}   { n := len(*pq); tmp := (*pq)[n-1]; *pq = (*pq)[:n-1]; return tmp }

// leetcode submit region end(Prohibit modification and deletion)
