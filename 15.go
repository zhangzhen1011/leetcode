import "sort"

// 双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums) // 一定要先排序
	n := len(nums)
	if n < 3 {
		return nil
	}

	var ret = make([][]int, 0)

	for i := 0; i < n; i++ { // 第一个
		if i > 0 && nums[i] == nums[i-1] { // 确定过一个之后再略过重复项
			continue
		}
		start, end := i+1, n-1
		for ; start < n; start++ { // 第二个
			if start > i+1 && nums[start] == nums[start-1] {
				continue
			}
			for start < end && nums[start]+nums[i]+nums[end] > 0 { // 第三个
				end--
			}
			if start == end {
				break
			}
			if nums[start]+nums[i]+nums[end] == 0 {
				ret = append(ret, []int{nums[i], nums[start], nums[end]})
			}
		}
	}

	return ret
}
