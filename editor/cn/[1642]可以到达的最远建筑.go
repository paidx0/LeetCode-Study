// 给你一个整数数组 heights ，表示建筑物的高度。另有一些砖块 bricks 和梯子 ladders 。
//
// 你从建筑物 0 开始旅程，不断向后面的建筑物移动，期间可能会用到砖块或梯子。
//
// 当从建筑物 i 移动到建筑物 i+1（下标 从 0 开始 ）时：
//
//
// 如果当前建筑物的高度 大于或等于 下一建筑物的高度，则不需要梯子或砖块
// 如果当前建筑的高度 小于 下一个建筑的高度，您可以使用 一架梯子 或 (h[i+1] - h[i]) 个砖块
// 如果以最佳方式使用给定的梯子和砖块，返回你可以到达的最远建筑物的下标（下标
// 从 0 开始 ）。
//
//
//
// 示例 1：
//
//
// 输入：heights = [4,2,7,6,9,14,12], bricks = 5, ladders = 1
// 输出：4
// 解释：从建筑物 0 出发，你可以按此方案完成旅程：
// - 不使用砖块或梯子到达建筑物 1 ，因为 4 >= 2
// - 使用 5 个砖块到达建筑物 2 。你必须使用砖块或梯子，因为 2 < 7
// - 不使用砖块或梯子到达建筑物 3 ，因为 7 >= 6
// - 使用唯一的梯子到达建筑物 4 。你必须使用砖块或梯子，因为 6 < 9
// 无法越过建筑物 4 ，因为没有更多砖块或梯子。
//
//
// 示例 2：
//
//
// 输入：heights = [4,12,2,7,3,18,20,3,19], bricks = 10, ladders = 2
// 输出：7
//
//
// 示例 3：
//
//
// 输入：heights = [14,3,19,3], bricks = 17, ladders = 0
// 输出：3
//
//
//
//
// 提示：
//
//
// 1 <= heights.length <= 10⁵
// 1 <= heights[i] <= 10⁶
// 0 <= bricks <= 10⁹
// 0 <= ladders <= heights.length
//
//
// Related Topics 贪心 数组 堆（优先队列） 👍 115 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
import (
	"container/heap"
)

type minheap struct {
	node []int
}

func (x minheap) Len() int           { return len(x.node) }
func (x minheap) Less(i, j int) bool { return x.node[i] < x.node[j] }
func (x minheap) Swap(i, j int)      { x.node[i], x.node[j] = x.node[j], x.node[i] }
func (x *minheap) Push(v interface{}) {
	x.node = append(x.node, v.(int))
}
func (x *minheap) Pop() interface{} {
	tmp := x.node[len(x.node)-1]
	x.node = x.node[:len(x.node)-1]
	return tmp
}

func (x *minheap) push(v int) {
	heap.Push(x, v)
}
func (x *minheap) pop() int {
	return heap.Pop(x).(int)
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	needLadders := new(minheap)
	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		// 当前楼比前一楼矮或等高，不需要梯子台阶
		if diff <= 0 {
			continue
		}
		// 否则需要
		if ladders > 0 {
			needLadders.push(diff)
			ladders--
		} else {
			// 楼梯的使用放在高度差最大的几个中
			// 如果当前高度比楼梯堆中大，则换出来用砖块，同时重新排序堆
			if needLadders.Len() > 0 && diff > needLadders.node[0] {
				diff, needLadders.node[0] = needLadders.node[0], diff
				heap.Fix(needLadders, 0)
			}
			bricks -= diff
			// 砖块不够了
			if bricks < 0 {
				return i - 1
			}
		}
	}
	return len(heights) - 1
}

// leetcode submit region end(Prohibit modification and deletion)
