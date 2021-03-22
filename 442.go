func findDuplicates(nums []int) []int {
	var result []int
	for _, val := range nums {
		if nums[abs(val)-1] < 0 {
			result = append(result, abs(val))
		} else {
			nums[abs(val)-1] = -nums[abs(val)-1]
		}
	}
	return result
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
