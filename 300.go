// dp[i] 代表区间 0..i, 0<=i<=n-1
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	//init
	dp := make([]int, len(nums))
	dp[0] = 1
	var ret = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		ret = max(ret, dp[i])
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
