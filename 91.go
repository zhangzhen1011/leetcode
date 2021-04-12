package main

// 问题转化：  有多少种切割，并且第一个数不为0
import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	count := 0
	dfs(s, &count)
	return count
}

// 超时
func dfs(left string, count *int) {
	if left == "" {
		*count = *count + 1
		return
	}
	if left[0] == '0' {
		return
	}
	n := len(left)
	for i := 1; i < n+1; i++ {
		intVal, _ := strconv.Atoi(left[:i])
		if intVal > 26 {
			break
		}
		dfs(left[i:], count)
	}
}

func main() {
	fmt.Println(numDecodings("226"))
}

func numDecodings(s string) int {
	dp := make([]int, len(s)+1)
	if s[0] == '0' {
		return 0
	}
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= len(s); i++ {
		cur, _ := strconv.Atoi(s[i-1 : i])
		if cur >= 1 && cur <= 9 {
			dp[i] += dp[i-1]
		}
		pre, _ := strconv.Atoi(s[i-2 : i])
		if pre >= 10 && pre <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
