// 有向图的环检测: 有向图、环
// 1. 构建图表示 2.dfs
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	marked := make([]bool, numCourses)
	onStack := make([]bool, numCourses)

	for _, vw := range prerequisites {
		adj[vw[0]] = append(adj[vw[0]], vw[1])
	}
	var hascycle bool
	for v := 0; v < numCourses; v++ {
		if hascycle {
			return false
		}
		if !marked[v] {
			dfs(adj, marked, onStack, v, &hascycle)
		}
	}

	return true
}

func dfs(adj [][]int, marked, onstack []bool, v int, hasCycle *bool) {
	onstack[v] = true
	marked[v] = true
	for _, w := range adj[v] {
		if *hasCycle {
			return
		} else if !marked[w] {
			dfs(adj, marked, onstack, w, hasCycle)
		} else if onstack[w] {
			*hasCycle = true
			return
		}
	}
	onstack[v] = false
}
