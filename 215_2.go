package main

// 找数组的第k大的数据
// FIXME
func findKthLargest(nums []int, k int) int {
	var lo, hi = 0, len(nums) - 1
	k = len(nums) - k
	for lo < hi {
		j := partition(nums, lo, hi)
		if j > k {
			hi = j - 1
		} else if j < k {
			lo = j + 1
		} else {
			return nums[j]
		}
	}
	return nums[lo]
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
	var a = []int{3, 2, 1, 5, 6, 4}
	println(findKthLargest(a, 2))
}
