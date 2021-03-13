// 如何分析
func findTargetSumWays(nums []int, S int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	if (total+S)%2 == 1 || S > total {
		return 0
	}

	target := (total + S) / 2
	var dp = make([]int, target+1)
	dp[0] = 1
	for _, v := range nums {
		for i := target; i >= v; i-- {
			dp[i] += dp[i-v]
		}
	}
	return dp[target]
}
