// ----------------v1 dp,超时
func tallestBillboard(rods []int) int {
	total := 0
	for _, val := range rods {
		total += val
	}
	dp := make([][]bool, 2501)
	for i := 0; i < 2501; i++ {
		dp[i] = make([]bool, 2501)
	}
	dp[0][0] = true //left = 0, right = 0
	var ret = 0
	for i := 0; i < len(rods); i++ {
		for left := total / 2; left >= 0; left-- {
			for right := total / 2; right >= 0; right-- {
				if dp[left][right] {
					if left+rods[i] < 2501 {
						dp[left+rods[i]][right] = true
					}
					if right+rods[i] < 2501 {
						dp[left][right+rods[i]] = true
					}
					if left+rods[i] == right {
						ret = max(ret, right)
					}
					if left == right+rods[i] {
						ret = max(ret, left)
					}
				}
			}
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

// 数组取两个集合s1，s2
// 都不为空
// 和相等, s1==s2, 2s1<=s, s1 <= s/2, 能取到的最大高度
// dp[left][right]: 记录所有可能

// -----------------v2 背包优化
// 不必枚举所有的left, right解空间，只需要针对diff值做计算，
// dp[diff]: 值为最大的left（同一个diff下的最大left），对diff空间染色
// diff最大值为5000（题目限定）
// 对于每个物品，可能加到左边，也可能加到右边
func tallestBillboard(rods []int) int {
	offset := 5000

	dp := make([]int, 2*offset+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0+offset] = 0 // 初始diff为0

	// 背包模板
	for i := 0; i < len(rods); i++ { // 对每个物品
		l := rods[i]
		tmp := make([]int, 2*offset+1)
		copy(tmp, dp)
		for diff := -offset; diff <= offset; diff++ { // 整个diff空间

			if tmp[diff+offset] == -1 { // 无效位置不传递状态
				continue
			}

			if diff+l <= offset {
				dp[diff+l+offset] = max(tmp[diff+offset]+l, tmp[diff+l+offset])
			}
			if diff-l+offset >= 0 {
				dp[diff-l+offset] = max(tmp[diff+offset], dp[diff-l+offset])
			}
		}
	}
	return dp[0+offset]
}
