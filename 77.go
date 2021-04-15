func combine(n int, k int) [][]int {
	var result = make([][]int, 0)
	dfs(n, 0, &result, []int{}, k)
	return result
}

// idx 为某一位的index, 表示当前的位置
func dfs(n int, idx int, result *[][]int, path []int, k int) {
	if idx > n {
		return
	}
	if len(path) == k {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}

	dfs(n, idx+1, result, append(path, idx+1), k)
	dfs(n, idx+1, result, path, k)
}

// 子集问题，只输出固定长度的子集
