// 有 n 个网络节点，标记为 1 到 n。
//
// 给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点，
// wi 是一个信号从源节点传递到目标节点的时间。
//
// 现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。
//
//
//
// 示例 1：
//
//
//
//
// 输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
// 输出：2
//
//
// 示例 2：
//
//
// 输入：times = [[1,2,1]], n = 2, k = 1
// 输出：1
//
//
// 示例 3：
//
//
// 输入：times = [[1,2,1]], n = 2, k = 2
// 输出：-1
//
//
//
//
// 提示：
//
//
// 1 <= k <= n <= 100
// 1 <= times.length <= 6000
// times[i].length == 3
// 1 <= ui, vi <= n
// ui != vi
// 0 <= wi <= 100
// 所有 (ui, vi) 对都 互不相同（即，不含重复边）
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 最短路 堆（优先队列） 👍 636 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n, k int) int {
	// 有向图
	g := make([][]edge, n+1)
	for _, t := range times {
		g[t[0]] = append(g[t[0]], edge{t[1], t[2]})
	}

	// 初始起点到各点的最短距离
	const inf int = math.MaxInt32 / 2
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = inf
	}
	dist[0] = 0

	// 起点
	dist[k] = 0
	h := &hp{pair{0, k}}
	for h.Len() > 0 {
		p := heap.Pop(h).(pair)
		x := p.to
		// 不更新
		if dist[x] < p.time {
			continue
		}
		// 更新
		for _, e := range g[x] {
			y := e.to
			if d := dist[x] + e.time; d < dist[y] {
				dist[y] = d
				heap.Push(h, pair{d, y})
			}
		}
	}

	var res int
	for _, d := range dist {
		// 存在不可达的点直接返回
		if d == inf {
			return -1
		}
		// 距离最远的就是花费时间
		res = max(res, d)
	}
	return res
}

// 有向边
type edge struct{ to, time int }

// 起点到点的距离
type pair struct{ time, to int }
type hp []pair

func (h hp) Len() int              { return len(h) }
func (h hp) Less(i, j int) bool    { return h[i].time < h[j].time }
func (h hp) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{})   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)
