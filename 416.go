// 等同于添加+/-，使和为0, 类似494
// sum(p) + sum(n) = total
// sum(p) - sum(n) = 0
// sum(p) + sum(n) + sum(p) - sum(n) = total + 0
// 即2sum(p) = total, 即能否找出子集，和为总和的一半
func canPartition(nums []int) bool {
	total := 0
	for _, v := range nums {
		total += v
	}
	if total%2 == 1 {
		return false
	}

	// init
	target := total / 2
	dp := make([]int, target+1) // dp[0..target]
	dp[0] = 1                   //初始值，表示取前0个物品的情况
	for n := 0; n < len(nums); n++ {
		for v := target; v >= nums[n]; v-- {
			dp[v] = dp[v] + dp[v-nums[n]]
		}
	}
	return dp[target] != 0 // dp[target] > 0, dp[target]值太大 会溢出。。TODO 改成bool值
}
