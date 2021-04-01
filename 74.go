func searchMatrix(matrix [][]int, target int) bool {
	var i int
	for i = 0; i < len(matrix); i++ {
		if target >= matrix[i][0] && target <= matrix[i][len(matrix[i])-1] {
			break
		}
	}
	if i == len(matrix) {
		return false
	}

	var lo, hi = 0, len(matrix[i]) - 1
	return binarySearch(a, lo, hi, target)
}

func binarySearch(a []int, lo, hi int, target int) bool {
	if lo > hi {
		return false
	}
	mid := (lo + hi) >> 1

	if a[mid] == target {
		return true
	} else if target < a[mid] {
		return binarySearch(a, lo, mid-1, target)
	} else {
		return binarySearch(a, mid+1, hi, target)
	}
}
