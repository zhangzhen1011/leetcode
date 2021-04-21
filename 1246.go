package main

import "fmt"

// 删除回文子数组
func minimumMoves(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	N := len(arr)
	arr = append([]int{0}, arr...)

	dp := make([][]int, N+2)
	for i := 0; i < N+2; i++ {
		dp[i] = make([]int, N+2)
	}

	for l := 1; l <= N; l++ {
		for i := 1; i+l-1 <= N; i++ {
			j := i + l - 1
			dp[i][j] = 1e10
			for k := i; k <= j; k++ {
				if arr[k] == arr[j] {
					dp[i][j] = min(dp[i][j], dp[i][k-1]+max(1, dp[k+1][j-1]))
				}
			}
		}
	}
	return dp[1][N]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// dp[i][j] : 表示arr[i..j] 之间的最少move数量
// dp[i][j], 在i....j， 如果s[k] == s[j]，那么i..j 可以分成i...k, k...j两部分, dp[i][j] = dp[i][k-1]+dp[k][j] = dp[i][k-1] + dp[k+1][j-1]

func main() {
	fmt.Println(minimumMoves([]int{1, 2, 3, 1}))
}
