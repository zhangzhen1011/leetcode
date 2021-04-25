// 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	col := len(matrix[0])
	i, j := 0, 0
	m, n := row-1, col-1
	var ret = []int{}
	for i <= m && j <= n {
		if j == n { // 一列
			for k := i; k <= m; k++ {
				ret = append(ret, matrix[k][j])
			}
		} else if i == m { // 一行
			for k := j; k <= n; k++ {
				ret = append(ret, matrix[i][k])
			}
		} else {
			for k := j; k < n; k++ {
				ret = append(ret, matrix[i][k])
			}
			for k := i; k < m; k++ {
				ret = append(ret, matrix[k][n])
			}
			for k := n; k > j; k-- {
				ret = append(ret, matrix[m][k])
			}
			for k := m; k > i; k-- {
				ret = append(ret, matrix[k][j])
			}
		}
		i++
		j++
		m--
		n--
	}
	return ret
}

// 矩阵：左上角坐标，和右下角坐标
