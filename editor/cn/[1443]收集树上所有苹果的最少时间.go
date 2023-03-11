// 给你一棵有 n 个节点的无向树，节点编号为 0 到 n-1 ，它们中有一些节点有苹果。通过树上的一条边，需要花费 1 秒钟。你从 节点 0 出发，请你返回最
// 少需要多少秒，可以收集到所有苹果，并回到节点 0 。
//
// 无向树的边由 edges 给出，其中 edges[i] = [fromi, toi] ，表示有一条边连接 from 和 toi 。除此以外，还有一个布尔数
// 组 hasApple ，其中 hasApple[i] = true 代表节点 i 有一个苹果，否则，节点 i 没有苹果。
//
//
//
// 示例 1：
//
//
//
//
// 输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,
// false,true,false,true,true,false]
// 输出：8
// 解释：上图展示了给定的树，其中红色节点表示有苹果。一个能收集到所有苹果的最优方案由绿色箭头表示。
//
//
// 示例 2：
//
//
//
//
// 输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,
// false,true,false,false,true,false]
// 输出：6
// 解释：上图展示了给定的树，其中红色节点表示有苹果。一个能收集到所有苹果的最优方案由绿色箭头表示。
//
//
// 示例 3：
//
//
// 输入：n = 7, edges = [[0,1],[0,2],[1,4],[1,5],[2,3],[2,6]], hasApple = [false,
// false,false,false,false,false,false]
// 输出：0
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10^5
// edges.length == n - 1
// edges[i].length == 2
// 0 <= ai < bi <= n - 1
// hasApple.length == n
//
//
// Related Topics 树 深度优先搜索 广度优先搜索 哈希表 👍 77 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// 子传父路径全标红，最后走一遍+2
// func minTime(n int, edges [][]int, hasApple []bool) int {
// 	var res int
// 	len_edges := len(edges)
// 	for i := len_edges - 1; i >= 0; i-- {
// 		if hasApple[edges[i][1]] == true {
// 			hasApple[edges[i][0]] = true
// 		}
// 	}
// 	for i := 0; i < len_edges; i++ {
// 		if hasApple[edges[i][1]] == true {
// 			res += 2
// 		}
// 	}
// 	return res
// }

func minTime(n int, edges [][]int, hasApple []bool) int {
	var res int
	g := make([][]int, n)   // 图
	vis := make([]bool, n)  // 标志走过的点，不能重复走返回
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		g[v[1]] = append(g[v[1]], v[0])
	}
	res = dfs(0, g, vis, hasApple, 0)
	return res
}

func dfs(curn int, g [][]int, vis []bool, hasApple []bool, cost int) int {
	if vis[curn] {
		return 0
	}
	vis[curn] = true
	childcost := 0
	for _, v := range g[curn] {
		childcost += dfs(v, g, vis, hasApple, 2)
	}
	// 子节点无苹果，且自身也无苹果，cost就是 0，没必要走这条路了
	if !hasApple[curn] && childcost == 0 {
		return 0
	}
	// 不管自身有苹果或者子节点有苹果，都需要走到这个节点，cost=2
	return childcost+cost
}

// leetcode submit region end(Prohibit modification and deletion)
