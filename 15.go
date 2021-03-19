import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	if n < 3 {
		return nil
	}

	var ret = make([][]int, 0)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, n-1
		for ; start < n; start++ {
			if start > i+1 && nums[start] == nums[start-1] {
				continue
			}
			for start < end && nums[start]+nums[i]+nums[end] > 0 {
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
