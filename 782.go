// hard
// 思路：先固定最左上角，然后固定其余部分
// 对任意单元, move grid[sr][sc]  to grid[0][0], with at most two switches (a col + a row switch)
// 使grid[0][c]有效，只用列交换
// 使grid[r][0]有效，只用行交换
// 此时，0行和0列已经满足
// 检查是否是有效棋盘，如果是，则满足
package main

import (
	"fmt"
)

var maxInt = 1<<63 - 1

func movesToChessboard(board [][]int) int {
	var res = maxInt
	if good(board) {
		return 0
	}
	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board); c++ {
			res = min(res, try(clone(board), r, c))
		}
	}
	if res == maxInt {
		return -1
	}
	return res
}

func try(grid [][]int, r, c int) int {
	var res int
	if r != 0 {
		res++
		switchRows(grid, r, 0)
	}
	if c != 0 {
		res++
		switchCols(grid, c, 0)
	}
	ret1 := setCol0Good(grid)
	ret2 := setRow0Good(grid)
	if ret1 != -1 && ret2 != -1 {
		res += ret1 + ret2
	}
	if good(grid) {
		return res
	}
	return maxInt
}

func clone(origin [][]int) [][]int {
	ret := [][]int{}
	for i := 0; i < len(origin); i++ {
		r := make([]int, len(origin))
		copy(r, origin[i])
		ret = append(ret, r)
	}
	return ret
}

func switchRows(grid [][]int, r1, r2 int) {
	grid[r1], grid[r2] = grid[r2], grid[r1]
}

func switchCols(grid [][]int, c1, c2 int) {
	for r := 0; r < len(grid); r++ {
		grid[r][c1], grid[r][c2] = grid[r][c2], grid[r][c1]
	}
}

func setCol0Good(grid [][]int) int {
	var comp1, comp2 = grid[0][0], grid[0][0] ^ 1
	var notgood1 []int
	for r := 1; r < len(grid); r += 2 {
		if grid[r][0] != comp2 {
			notgood1 = append(notgood1, r)
		}
	}
	var notgood2 []int
	for r := 2; r < len(grid); r += 2 {
		if grid[r][0] != comp1 {
			notgood2 = append(notgood2, r)
		}
	}

	if len(notgood1) != len(notgood2) {
		return -1
	}
	for i := 0; i < len(notgood1); i++ {
		switchRows(grid, notgood1[i], notgood2[i])
	}
	return len(notgood1)
}

func setRow0Good(grid [][]int) int {
	var comp1, comp2 = grid[0][0], grid[0][0] ^ 1
	var notgood1 []int
	for c := 1; c < len(grid); c += 2 {
		if grid[0][c] != comp2 {
			notgood1 = append(notgood1, c)
		}
	}
	var notgood2 []int
	for c := 2; c < len(grid); c += 2 {
		if grid[0][c] != comp1 {
			notgood2 = append(notgood2, c)
		}
	}

	if len(notgood1) != len(notgood2) {
		return -1
	}
	for i := 0; i < len(notgood1); i++ {
		switchCols(grid, notgood1[i], notgood2[i])
	}
	return len(notgood1)
}

func good(grid [][]int) bool {
	if len(grid) == 0 || len(grid) == 1 {
		return true
	}

	for c := 1; c < len(grid); c++ {
		if grid[0][c] == grid[0][c-1] {
			return false
		}
	}

	for r := 1; r < len(grid); r++ {
		if grid[r][0] == grid[r-1][0] {
			return false
		}
		for c := 1; c < len(grid); c++ {
			if grid[r][c] == grid[r][c-1] || grid[r][c] == grid[r-1][c] {
				return false
			}
		}
	}
	return true
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func main() {
	fmt.Println(movesToChessboard([][]int{{1, 1, 0}, {0, 0, 1}, {0, 0, 1}}))
}
