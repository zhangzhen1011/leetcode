// dp[i],  以i为结尾的最大和
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	var ret int = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > ret {
			ret = dp[i]
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
