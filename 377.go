// 完全背包问题
// 每个物品可以无限选择
// dp数组记录
// 如果求排列，外层循环是背包
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for j := 0; j < target+1; j++ {
		for i := 0; i < len(nums); i++ {
			if j >= nums[i] {
				dp[j] = dp[j] + dp[j-nums[i]]
			}
		}
	}
	return dp[target]
}
