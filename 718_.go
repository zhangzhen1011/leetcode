// 最长公共子数组
// 1. dp
func findLength(A []int, B []int) int {
	dp := make([][]int, len(A)+1)
	for i := 0; i <= len(A); i++ {
		dp[i] = make([]int, len(B)+1)
	}

	var result int

	for i := 1; i <= len(A); i++ {
		for j := 1; j <= len(B); j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			result = max(result, dp[i][j])
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// TODO
// 2. 滑动窗口（两个数组）

// 3. 二分查找+hash
