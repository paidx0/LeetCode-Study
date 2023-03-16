// 树可以看成是一个连通且 无环 的 无向 图。
//
// 给定往一棵 n 个节点 (节点值 1～n) 的树中添加一条边后的图。添加的边的两个顶点包含在 1 到 n 中间，且这条附加的边不属于树中已存在的边。图的信
// 息记录于长度为 n 的二维数组 edges ，edges[i] = [ai, bi] 表示图中在 ai 和 bi 之间存在一条边。
//
// 请找出一条可以删去的边，删除后可使得剩余部分是一个有着 n 个节点的树。如果有多个答案，则返回数组 edges 中最后出现的边。
//
//
//
// 示例 1：
//
//
//
//
// 输入: edges = [[1,2], [1,3], [2,3]]
// 输出: [2,3]
//
//
// 示例 2：
//
//
//
//
// 输入: edges = [[1,2], [2,3], [3,4], [1,4], [1,5]]
// 输出: [1,4]
//
//
//
//
// 提示:
//
//
// n == edges.length
// 3 <= n <= 1000
// edges[i].length == 2
// 1 <= ai < bi <= edges.length
// ai != bi
// edges 中无重复元素
// 给定的图是连通的
//
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 图 👍 542 👎 0

// leetcode submit region begin(Prohibit modification and deletion)

type Unionfind struct {
	list []int
}

// 并查集，树中的边数应该比节点数少一个，初始图指向自己
func NewUnionfind(n int) *Unionfind {
	l := make([]int, n)
	for i := 1; i < n; i++ {
		l[i] = i
	}
	return &Unionfind{
		list: l,
	}
}
// 往上找到父节点
func (receiver Unionfind) find(p int) int {
	if receiver.list[p] != p {
		receiver.list[p] = receiver.find(receiver.list[p])
	}
	return receiver.list[p]
}
// 共父节点返回false
func (receiver Unionfind) union(p, q int) bool {
	from, to := receiver.find(p), receiver.find(q)
	if from == to {
		return false
	}
	receiver.list[from] = to
	return true
}

func findRedundantConnection(edges [][]int) []int {
	parent := NewUnionfind(len(edges) + 1)
	for _, e := range edges {
		if !parent.union(e[0], e[1]) {
			return e
		}
	}
	return []int{}
}

// leetcode submit region end(Prohibit modification and deletion)
