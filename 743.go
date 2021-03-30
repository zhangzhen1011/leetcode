// 网络延迟，dfs
// TODO，超时边缘，优化
func networkDelayTime(times [][]int, n int, k int) int {
	// 构建
	g := make([][]node, n+1)
	for _, item := range times {
		n := node{item[1], item[2]}
		g[item[0]] = append(g[item[0]], n)
	}

	// marked := make([]bool, n+1)
	// onstack := make([]bool, n+1) // 递归栈记录
	pathSum := make([]int, n+1)
	for idx, _ := range pathSum {
		pathSum[idx] = 1e10
	}

	dfs(g, k, 0, pathSum)
	var max int
	for idx, val := range pathSum {
		if idx == 0 {
			continue
		}
		if val == 1e10 {
			return -1
		}
		if val > max {
			max = val
		}
	}
	return max
}

type node struct {
	val int
	wei int
}

func dfs(g [][]node, v int, total int, pathsum []int) {
	if total >= pathsum[v] {
		return
	}
	if total < pathsum[v] {
		pathsum[v] = total
	}
	for idx, n := range g[v] {
		dfs(g, n.val, total+n.wei, pathsum)
	}
}

