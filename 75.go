func sortColors(nums []int) {
	sort(nums)
}

func partition(nums []int, lo, hi int) {
	var i, j := lo+1, hi
	for i < j {
		for i <= hi {
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

		if i > j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	if j > lo {
		nums[lo], nums[j] = nums[j], nums[lo]
	}
	return j
}

func sort(nums []int) {
	p := partition(nums, 0, len(nums)-1)
	sort(nums, 0, p-1)
	sort(nums, p+1, len(nums)-1)
}
