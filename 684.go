// 无向图，找环
// TODO 11% beat
func findRedundantConnection(edges [][]int) []int {
	g := make([][]int, len(edges)+1)
	for _, val := range edges {
		g[val[0]] = append(g[val[0]], val[1])
		g[val[1]] = append(g[val[1]], val[0]) //NOTE ： 无向图addedge两次
	}
	marked := make([]bool, len(edges)+1)
	var ret = make([]bool, len(edges)+1)
	var edgeTo = make([]int, len(edges)+1)
	for v := 1; v < len(g); v++ {
		dfs(g, v, -1, marked, ret, edgeTo)
	}
	for i := len(edges) - 1; i >= 0; i-- {
		if ret[edges[i][0]] && ret[edges[i][1]] {
			return edges[i]
		}
	}
	return nil
}

func dfs(g [][]int, v int, u int, marked []bool, ret []bool, edgeTo []int) {
	marked[v] = true
	for _, w := range g[v] {
		if ret[0] {
			return
		}
		if !marked[w] {
			edgeTo[w] = v
			dfs(g, w, v, marked, ret, edgeTo)
		} else if w != u {
			ret[0] = true
			for x := v; x != w; x = edgeTo[x] {
				ret[x] = true
			}
			ret[w] = true
			ret[v] = true
		}
	}
}
