var MinValue int = -1e10

func maxProfit(prices []int) int {
	// 没必要记录每一天，只要记住今天，和前一天的状态就行
	//dp := make([][]int, len(prices)+1)
	//for i := 0; i <= len(prices); i++ {
	//	dp[i] = make([]int, 4)
	//	dp[i][0] = MinValue
	//	dp[i][1] = MinValue
	//	dp[i][2] = MinValue
	//	dp[i][3] = MinValue
	//}
	var hold1, sold1, hold2, sold2 = MinValue, MinValue, MinValue, MinValue

	for i := 1; i <= len(prices); i++ {
		p := prices[i-1]
		hold1_tmp := hold1
		sold1_tmp := sold1
		hold2_tmp := hold2
		sold2_tmp := sold2
		hold1 = max(-p, hold1_tmp)
		sold1 = max(hold1_tmp+p, sold1_tmp)
		hold2 = max(sold1_tmp-p, hold2_tmp)
		sold2 = max(hold2_tmp+p, sold2_tmp)
	}

	ret := max(sold1, sold2)
	if ret < 0 {
		return 0
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// 顺着时间线，每天4种状态，持有或买进第一只，卖出第一只，持有或买进第二只，卖出第二只
// dp[i][0..4]
// dp[i][0] = max(-p, dp[i-1][0])
// dp[i][1] = max(dp[i-1][0]+p, dp[i-1][1])
// dp[i][2] = max(dp[i-1][1]-p, dp[i-1][2])
// dp[i][3] = max(dp[i-1][2]+p, dp[i-1][3])
