func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i][0] = max(dp[i-1][1], dp[i-1][0])
		dp[i][1] = dp[i-1][0] + nums[i]
	}
	return max(dp[len(nums)-1][0], dp[len(nums)-1][1])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// dp[i][0], dp[i][1]
