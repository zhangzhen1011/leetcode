func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(rob1(nums[:len(nums)-1]), rob1(nums[1:]))
}

func rob1(nums []int) int {

	dp := make([]int, len(nums)+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}
	return dp[len(nums)]
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//dp[i]: [1..i] 区间，且选择第i个位置
// dp[i] : 两种情况取较大，dp[i-1], dp[i-2]+nums[i]
