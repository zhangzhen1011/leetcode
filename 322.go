func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0 // init，初始背包空间为0时，需要0个硬币
	for i := 1; i < len(dp); i++ {
		dp[i] = 1e10
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j < amount+1 && j >= coins[i]; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == 1e10 {
		return -1
	}
	return dp[amount]
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
