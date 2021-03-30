// 拓扑排序，在207的基础上修改
func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := make([][]int, numCourses)
	marked := make([]bool, numCourses)
	onStack := make([]bool, numCourses)
	postOrder := []int{}

	for _, vw := range prerequisites {
		adj[vw[0]] = append(adj[vw[0]], vw[1])
	}
	var hascycle bool
	for v := 0; v < numCourses; v++ {
		if hascycle {
			return nil
		}
		if !marked[v] {
			dfs(adj, &postOrder, marked, onStack, v, &hascycle)
		}
	}

	return postOrder
}

func dfs(adj [][]int, order *[]int, marked, onstack []bool, v int, hasCycle *bool) {
	onstack[v] = true
	marked[v] = true
	for _, w := range adj[v] {
		if *hasCycle {
			return
		} else if !marked[w] {
			dfs(adj, order, marked, onstack, w, hasCycle)
		} else if onstack[w] {
			*hasCycle = true
			return
		}
	}
	*order = append(*order, v)
	onstack[v] = false
}
