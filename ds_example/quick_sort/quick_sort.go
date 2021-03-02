package main

import "fmt"

// a：数组，lo，hi，低地址，高地址，左闭右开
// 返回切分位置
func partition(a []int, lo, hi int) int {
	i := lo + 1
	j := hi - 1
	for i <= j {
		for i < hi {
			if a[i] > a[lo] {
				break
			}
			i++
		}
		for j > lo {
			if a[j] < a[lo] {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j // 要返回j, 因为交换前可能i和j相交，a[j] < a[lo], a[i]>a[lo], 如果和a[i] 交换，又会产生逆序对
}

// 3, 1,1,4,2,5,5,6
// 3, 1,1,2,4,5,5,6

func QuickSort(a []int, lo, hi int) {
	if lo >= hi {
		return
	}

	var j = partition(a, lo, hi)
	QuickSort(a, lo, j)
	QuickSort(a, j+1, hi)
}

func main() {
	a := []int{1, 4, 3, 5, 2, 1}
	QuickSort(a, 0, len(a))
	fmt.Println(a)

	a = []int{0, 1}
	QuickSort(a, 0, len(a))
	fmt.Println(a)
}
