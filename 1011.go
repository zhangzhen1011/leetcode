// 看着像是切分数组到D个子区间，使每个部分的和都小于sum， 要求sum尽可能小
// 类似410
// 二分法
func shipWithinDays(weights []int, D int) int {
	var left, right int
	left = -1e10
	for _, val := range weights {
		left = max(left, val)
		right += val
	}

	var capacity int
	for left < right {
		capacity = (left + right) / 2
		if checkOk(weights, capacity, D) { // capacity 足够
			right = capacity
		} else {
			left = capacity + 1
		}
	}
	return left
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func checkOk(nums []int, capacity int, D int) bool {
	count := 0
	for i := 0; i < len(nums); {
		j := i
		sum := 0
		for j < len(nums) {
			sum += nums[j]
			if sum > capacity {
				break
			}
			j++
		}
		count++
		if count > D {
			return false
		}
		i = j
	}
	return true
}
