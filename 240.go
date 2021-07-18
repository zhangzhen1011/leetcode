// 搜索二维矩阵
// 左右有序，上下有序
func searchMatrix(matrix [][]int, target int) bool {
	var ret bool
	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] > target {
			break
		}
		if matrix[i][len(matrix[i])-1] < target {
			continue
		}
		ret = ret || binary(matrix[i], 0, len(matrix[i])-1, target)
	}
	return ret
}

func binary(nums []int, lo, hi int, target int) bool {
	if lo <= hi {
		if target > nums[hi] || target < nums[lo] {
			return false
		}
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] > target {
			return binary(nums, lo, mid-1, target)
		}
		return binary(nums, mid+1, hi, target)
	}
	return false
}
