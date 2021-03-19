func generateParenthesis(n int) []string {
	var result = &[]string{}
	helper(n, n, "", result)
	return *result
}

// 参数1. 保存结果result 2. 剩余遍历范围 3. path保存路径
func helper(l, r int, path string, result *[]string) {
	// 下探过程保证 l <= r
	if l == 0 && r == 0 {
		*result = append(*result, path)
		return
	}
	if l < 0 || r < 0 {
		return
	}
	if l == 0 {
		helper(l, r-1, path+")", result)
	} else {
		if r > l {
			helper(l-1, r, path+"(", result)
			helper(l, r-1, path+")", result)
		} else {
			helper(l-1, r, path+"(", result)
		}
	}
}
