import (
	//"fmt"
	"sort"
)

// ---------------------------------v1 bf
// 搜索超时
func splitArraySameAverage(nums []int) bool {
	// 转化
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = nums[i]*len(nums) - total
	}

	var ret bool
	sort.Ints(nums)
	//dfs(nums, []int{}, 0, &ret)
	dfs2(nums, 0, 0, 0, &ret)
	return ret
}

// dfs 查询所有可能子集，是否有和为0的子集
func dfs(left []int, path []int, total int, result *bool) {
	if total == 0 {
		if len(path) > 0 && len(left) > 0 {
			*result = true
			return
		}
	}
	if len(left) == 0 {
		return
	}

	for i := 0; i < len(left); i++ {
		dfs(left[i+1:], append(path, left[i]), total+left[i], result)
	}
}

func dfs2(nums []int, curNum int, total int, idx int, result *bool) {
	if curNum != len(nums) && total == 0 && curNum != 0 {
		*result = true
		return
	}
	if idx == len(nums) {
		return
	}
	if total > 0 && nums[idx] > 0 {
		return
	}

	// 选择
	dfs2(nums, curNum+1, total+nums[idx], idx+1, result)
	// 不选择
	var i = idx + 1
	for i < len(nums) && nums[i] == nums[idx] {
		i++
	}
	dfs2(nums, curNum, total, i, result)
}

// -------------------------------------v2 bf
// 另一种搜索策略，用例都过了 勉强ac
import (
	"sort"
)

func splitArraySameAverage(nums []int) bool {
	sort.Ints(nums)
	n := len(nums)
	var total int
	for _, val := range nums {
		total += val
	}
	for num := 1; num < n; num++ {
		if total*num%n != 0 { // num*total = sum*n
			continue
		}
		sum := total * num / n
		if dfs(nums, num, sum, 0) {
			return true
		}
	}
	return false
}

func dfs(nums []int, num, sum, idx int) bool {
	if num == 0 && sum == 0 {
		return true
	}
	if num <= 0 || sum <= 0 {
		return false
	}

	if idx == len(nums) {
		return false
	}

	// 选择
	if dfs(nums, num-1, sum-nums[idx], idx+1) {
		return true
	}

	var i = idx + 1
	for i < len(nums) && nums[i] == nums[idx] {
		i++
	}
	// 不选择
	if dfs(nums, num, sum, i) {
		return true
	}
	return false
}


//  -------------------v3 dp 背包
import "sort"
// dp[sum][num] ||= dp[num-nums[i]][num-1]
func splitArraySameAverage(nums []int) bool {
	sort.Ints(nums)
	n := len(nums)
	var total int
	for _, val := range nums {
		total += val
	}

	dp := make([][]bool, total+1)
	for i := 0; i <= total; i++ {
		dp[i] = make([]bool, n/2+2)
	}
	dp[0][0] = true // 可能会有sum为0，num为0

	var curSum = 0
	for _, a := range nums {
		curSum += a
		for sum := curSum; sum >= a; sum-- {
			for num := n/2 + 1; num >= 1; num-- { // 一个集合多余一半，另一个集合必定少于一半，求一个就行
				dp[sum][num] = dp[sum-a][num-1] || dp[sum][num]
				if num != n && dp[sum][num] && n*sum == total*num {
					return true
				}
			}
		}
	}
	return false
}

// --------------v4 背包优化,  二维变成一维dp
// dp[sum][num] ||= dp[num-nums[i]][num-1]
func splitArraySameAverage(nums []int) bool {
	if len(nums) == 1 {
		return false
	}
	//sort.Ints(nums)
	n := len(nums)
	var total int
	for _, val := range nums {
		total += val
	}

	dp := make([]uint32, total+1) // 0b10010, 位运算表示， 这个例子表示，sum下num=1时，为true， num=4也为true
	dp[0] = 0b1

	var curSum = 0
	for _, a := range nums {
		curSum += a
		for sum := curSum; sum >= a; sum-- {
			dp[sum] |= dp[sum-a] << 1 // 状态移位，例如sum-a时，0b1010, 1，3位为true，那么sum时，num+1， 即2，4位为1，即0b10100
			if n*sum%total != 0 {
				continue
			}
			num := n * sum / total
			if num != 0 && num != n && dp[sum]&(1<<num) != 0 {
				return true
			}
		}
	}
	return false
}
