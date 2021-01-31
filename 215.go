package main

// 找数组的第k大的数据
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, lo, hi, index int) int {
	j := partition(a, lo, hi)
	if j == index {
		return a[j]
	}
	if j < index {
		return quickSelect(a, j+1, hi, index)
	}
	return quickSelect(a, lo, j-1, index)
}

func partition(a []int, lo, hi int) int {
	x := a[hi]
	i := lo - 1
	for j := lo; j < hi; j++ {
		if a[j] <= x {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}

	a[i+1], a[hi] = a[hi], a[i+1]
	return i + 1
}

func main() {
	var a = []int{3, 2, 3, 1}
	println(findKthLargest(a, 2))
}
