// 公共超序列, 难点构造字符串
func shortestCommonSupersequence(str1 string, str2 string) string {

	m := len(str1)
	n := len(str2)
	str1 = "#" + str1
	str2 = "#" + str2

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	path := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		path[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		path[i][0] = 2
	}
	for j := 1; j <= n; j++ {
		path[0][j] = 3
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i] == str2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				path[i][j] = 1
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + 1

				// 记录path
				if dp[i][j] == dp[i-1][j]+1 {
					path[i][j] = 2
				} else {
					path[i][j] = 3
				}
			}
		}
	}
	// dp[m][n]

	var ret string
	i, j := m, n
	for i > 0 || j > 0 {
		if path[i][j] == 1 {
			ret = string(str1[i]) + ret
			i--
			j--
		} else if path[i][j] == 2 {
			ret = string(str1[i]) + ret
			i--
		} else {
			ret = string(str2[j]) + ret
			j--
		}
	}
	return ret
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// 双序列型dp
// dp[i][j] : str1[0..i], str2[0..j]
// 先计算最短超序列的长度，然后再逆序构造路径
