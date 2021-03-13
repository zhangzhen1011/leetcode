func nextGreaterElements(nums []int) []int {
	n := len(nums)
	var res = make([]int, n)
	for k, _ := range res {
		res[k] = -1
	}
	stack := []int{}
	for i := 0; i < n*2; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			res[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return res
}
