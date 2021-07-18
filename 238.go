// 左右累乘
func productExceptSelf(nums []int) []int {
	var ret = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ret[i] = 1
	}
	left, right := 1, 1
	for i := 0; i < len(nums); i++ {
		ret[i] *= left
		left *= nums[i]

		ret[len(nums)-1-i] *= right
		right *= nums[len(nums)-1-i]
	}
	return ret
}
