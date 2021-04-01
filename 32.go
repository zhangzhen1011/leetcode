// 32 hard
// 1. BF， n(n+1)/2个子串，遍历最长子串
// time out
func longestValidParentheses(s string) int {
	var i, j = 0, 0
	var result int
	for i = 0; i < len(s); i++ {
		for j = i; j < len(s); j++ {
			if valid(s[i : j+1]) {
				result = max(result, j-i+1)
			}
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

// 判断有效括号
func valid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	bytes := []byte(s)
	if len(bytes) == 0 {
		return true
	}
	stack := []byte{}
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == '(' {
			stack = append(stack, bytes[i])
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

// 2. DP
// 最优子结构
func longestValidParenthesesDP(s string) int {
	if len(s) < 2 {
		return 0
	}
	dp := make([]int, len(s))
	if s[0] == '(' && s[1] == ')' {
		dp[1] = 2
	}

	var result = dp[1]
	for i := 2; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				dp[i] = dp[i-2] + 2
			} else {
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = dp[i-1] + 2
					if i-dp[i-1]-2 >= 0 { // NOTE,  这里容易漏
						dp[i] += dp[i-dp[i-1]-2]
					}
				}
			}
			result = max(dp[i], result)
		}
	}

	return result
}

// 3. 栈
// 4. 贪心，平衡, TODO
