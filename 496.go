// 单调栈
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := []int{}
	var m = map[int]int{}

	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			m[nums2[i]] = -1
		} else {
			m[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	var ret []int
	for _, v := range nums1 {
		ret = append(ret, m[v])
	}

	return ret
}

// 补充
