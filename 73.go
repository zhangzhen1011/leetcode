func setZeroes(matrix [][]int) {
	var a, b = map[int]int{}, map[int]int{}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				a[i] = 1
				b[j] = 1
			}
		}
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if _, ok := a[i]; ok {
				matrix[i][j] = 0
			}
			if _, ok := b[j]; ok {
				matrix[i][j] = 0
			}
		}
	}
}
