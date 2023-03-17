// 给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。
//
//
//
// 示例 1:
//
//
// 输入: nums = [1,1,1,2,2,3], k = 2
// 输出: [1,2]
//
//
// 示例 2:
//
//
// 输入: nums = [1], k = 1
// 输出: [1]
//
//
//
// 提示：
//
//
// 1 <= nums.length <= 10⁵
// k 的取值范围是 [1, 数组中不相同的元素的个数]
// 题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
//
//
//
//
// 进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
//
// Related Topics 数组 哈希表 分治 桶排序 计数 快速选择 排序 堆（优先队列） 👍 1496 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
import (
	"container/heap"
)

func topKFrequent(nums []int, k int) []int {
	src := make(map[int]int, 0)
	for _, v := range nums {
		src[v]++
	}
	hp := maxheap{}
	for k, v := range src {
		hp.push(node{key: k, cnt: v})
	}
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, hp.pop())
	}
	return res
}

type node struct {
	key int
	cnt int
}
type maxheap struct {
	node []node
}

func (x maxheap) Len() int           { return len(x.node) }
func (x maxheap) Less(i, j int) bool { return x.node[i].cnt > x.node[j].cnt }
func (x maxheap) Swap(i, j int)      { x.node[i], x.node[j] = x.node[j], x.node[i] }
func (x *maxheap) Push(v interface{}) {
	x.node = append(x.node, v.(node))
}
func (x *maxheap) Pop() interface{} {
	tmp := x.node[len(x.node)-1]
	x.node = x.node[:len(x.node)-1]
	return tmp.key
}

func (x *maxheap) push(v node) {
	heap.Push(x, v)
}
func (x *maxheap) pop() int {
	return heap.Pop(x).(int)
}

// leetcode submit region end(Prohibit modification and deletion)
