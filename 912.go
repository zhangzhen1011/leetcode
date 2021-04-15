func sortArray(nums []int) []int {
	qsort(nums, 0, len(nums))
	return nums
}

func qsort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	j := partition(nums, lo, hi)
	qsort(nums, lo, j)
	qsort(nums, j+1, hi)
}

func partition(nums []int, lo, hi int) int {
	i, j := lo+1, hi-1
	for i <= j { // 边界NOTICE
		for i < hi {
			if nums[i] > nums[lo] {
				break
			}
			i++
		}

		for j > lo {
			if nums[j] < nums[lo] {
				break
			}
			j--
		}
		if i >= j { // 边界NOTICE
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	nums[lo], nums[j] = nums[j], nums[lo]
	return j
}
