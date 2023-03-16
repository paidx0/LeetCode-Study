// 在本问题中，有根树指满足以下条件的 有向 图。该树只有一个根节点，所有其他节点都是该根节点的后继。该树除了根节点之外的每一个节点都有且只有一个父节点，而根节
// 点没有父节点。
//
// 输入一个有向图，该图由一个有着 n 个节点（节点值不重复，从 1 到 n）的树及一条附加的有向边构成。附加的边包含在 1 到 n 中的两个不同顶点间，这条
// 附加的边不属于树中已存在的边。
//
// 结果图是一个以边组成的二维数组 edges 。 每个元素是一对 [ui, vi]，用以表示 有向 图中连接顶点 ui 和顶点 vi 的边，其中 ui 是
// vi 的一个父节点。
//
// 返回一条能删除的边，使得剩下的图是有 n 个节点的有根树。若有多个答案，返回最后出现在给定二维数组的答案。
//
//
//
// 示例 1：
//
//
// 输入：edges = [[1,2],[1,3],[2,3]]
// 输出：[2,3]
//
//
// 示例 2：
//
//
// 输入：edges = [[1,2],[2,3],[3,4],[4,1],[1,5]]
// 输出：[4,1]
//
//
//
//
// 提示：
//
//
// n == edges.length
// 3 <= n <= 1000
// edges[i].length == 2
// 1 <= ui, vi <= n
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 348 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

// 有向图，入度都为1，如果入度为2，说明有边多余
func findRedundantDirectedConnection(edges [][]int) (redundantEdge []int) {
	n := len(edges)
	uf := NewUnionfind(n + 1)
	// 存父节点，初始指向自己
	parent := make([]int, n+1)
	for i := range parent {
		parent[i] = i
	}
	var conflictEdge, cycleEdge []int
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		if parent[to] != to {
			// 说明入度为2出现
			conflictEdge = edge
		} else {
			// 存父节点
			parent[to] = from
			if uf.find(from) == uf.find(to) {
				cycleEdge = edge
			} else {
				uf.union(from, to)
			}
		}
	}
	// 若不存在一个节点有两个父节点的情况，则可能附加边在根节点
	if conflictEdge == nil {
		return cycleEdge
	}
	// conflictEdge[1]点的入度为2，其中之一与其构成附加的边
	if cycleEdge != nil {
		return []int{parent[conflictEdge[1]], conflictEdge[1]}
	}
	return conflictEdge
}

type Unionfind struct {
	list []int
}

func NewUnionfind(n int) *Unionfind {
	l := make([]int, n)
	for i := 1; i < n; i++ {
		l[i] = i
	}
	return &Unionfind{
		list: l,
	}
}
func (receiver Unionfind) find(p int) int {
	if receiver.list[p] != p {
		receiver.list[p] = receiver.find(receiver.list[p])
	}
	return receiver.list[p]
}
func (receiver Unionfind) union(p, q int) bool {
	from, to := receiver.find(p), receiver.find(q)
	if from == to {
		return false
	}
	receiver.list[from] = to
	return true
}

// leetcode submit region end(Prohibit modification and deletion)
