func increasingTriplet(nums []int) bool {
	// dp[i], 代表0..i
	dp := make([]int, len(nums)+1)

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if dp[i] >= 2 {
			return true
		}
	}
	return false
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
