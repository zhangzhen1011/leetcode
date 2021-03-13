func nextPermutation(nums []int) {
	// 找两个数
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}

	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse(&nums, i+1, len(nums)-1)
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}
