// 有 n 个城市通过一些航班连接。给你一个数组 flights ，其中 flights[i] = [fromi, toi, pricei] ，表示该航班都从城
// 市 fromi 开始，以价格 pricei 抵达 toi。
//
// 现在给定所有的城市和航班，以及出发城市 src 和目的地 dst，你的任务是找到出一条最多经过 k 站中转的路线，使得从 src 到 dst 的 价格最便
// 宜 ，并返回该价格。 如果不存在这样的路线，则输出 -1。
//
//
//
// 示例 1：
//
//
// 输入:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
// src = 0, dst = 2, k = 1
// 输出: 200
// 解释:
// 城市航班图如下
//
//
// 从城市 0 到城市 2 在 1 站中转以内的最便宜价格是 200，如图中红色所示。
//
// 示例 2：
//
//
// 输入:
// n = 3, edges = [[0,1,100],[1,2,100],[0,2,500]]
// src = 0, dst = 2, k = 0
// 输出: 500
// 解释:
// 城市航班图如下
//
//
// 从城市 0 到城市 2 在 0 站中转以内的最便宜价格是 500，如图中蓝色所示。
//
//
//
// 提示：
//
//
// 1 <= n <= 100
// 0 <= flights.length <= (n * (n - 1) / 2)
// flights[i].length == 3
// 0 <= fromi, toi < n
// fromi != toi
// 1 <= pricei <= 10⁴
// 航班没有重复，且不存在自环
// 0 <= src, dst, k < n
// src != dst
//
//
// Related Topics 深度优先搜索 广度优先搜索 图 动态规划 最短路 堆（优先队列） 👍 560 👎 0

// leetcode submit region begin(Prohibit modification and deletion)
import (
	"math"
)

// 动态规划，最多 k+1 次乘坐
// f[k+1][dst] = min(f[k][j]+flights[j][dst])
// 多次乘坐最小值 min(f[1][dst]、f[2][dst]、f[3][dst]。。。。。f[k+1][dst])
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	const inf = math.MaxInt32 / 2
	f := make([]int, n)
	for i := range f {
		f[i] = inf
	}
	// 起点
	f[src] = 0
	var res = inf
	// 更新 k+1 次
	for t := 1; t <= k+1; t++ {
		g := make([]int, n)
		for i := range g {
			g[i] = inf
		}
		for _, time := range flights {
			from, to, cost := time[0], time[1], time[2]
			g[to] = min(g[to], f[from]+cost)
		}
		f = g
		res = min(res, g[dst])
	}
	if res == inf {
		return -1
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// leetcode submit region end(Prohibit modification and deletion)
