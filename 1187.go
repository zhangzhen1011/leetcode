package main

import (
	"fmt"
	"sort"
)

var MaxValue int = 1e10

// dp[i][k] i: arr1[1..n], 第二维k表示进行了k次操作。   值： k次操作过程中的最小值。
// dp[i][k] 从i-1传导过来，有两个选择: dp[i-1][k], dp[i-1][k-1]
func makeArrayIncreasing(arr1 []int, arr2 []int) int {
	sort.Ints(arr2)
	//m := map[int]int{}
	//for _, v := range arr1 {
	//	if val, ok := find(arr2, v); ok {
	//		m[v] = val
	//	}
	//}
	//for _, v := range arr2 {
	//	if val, ok = find(arr2, v); ok {
	//		m[v] = val
	//	}
	//}
	//if val, ok := find(arr2, -MaxValue); ok {
	//	m[-MaxValue] = val
	//}
	dp := make([][]int, len(arr1)+1)
	for i := 0; i <= len(arr1); i++ {
		dp[i] = make([]int, len(arr1)+1)

		for j := 0; j <= len(arr1); j++ {
			dp[i][j] = MaxValue
		}
	}
	dp[0][0] = -MaxValue
	for i := 1; i <= len(arr1); i++ {
		for k := 0; k <= i; k++ {
			if dp[i-1][k] < arr1[i-1] {
				dp[i][k] = min(arr1[i-1], dp[i][k])
			}

			if k > 0 {
				val, ok := bifind(arr2, 0, len(arr2), dp[i-1][k-1])
				if ok {
					dp[i][k] = min(dp[i][k], val)
				}
			}
		}
	}
	var ret int = -1
	for k, v := range dp[len(arr1)] {
		if v != MaxValue {
			ret = k
			break
		}
	}
	return ret
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func find(a []int, target int) (int, bool) {
	for _, val := range a {
		if val > target {
			return val, true
		}
	}
	return 0, false
}

func bifind(a []int, lo, hi int, target int) (int, bool) {
	if lo >= hi {
		fmt.Println("---", lo, hi)
		if hi > len(a)-1 {
			return 0, false
		}
		if a[hi] <= target {
			return 0, false
		}
		return a[hi], true
	}
	mid := (lo + hi) / 2
	if a[mid] > target {
		return bifind(a, lo, mid, target)
	} else {
		return bifind(a, mid+1, hi, target)
	}
}

func main() {
	a := []int{1, 2, 3, 3, 4, 4}
	fmt.Println(bifind(a, 0, len(a)-1, 5))
	fmt.Println(bifind(a, 0, len(a)-1, 4))
	fmt.Println(bifind(a, 0, len(a)-1, 3))
}
