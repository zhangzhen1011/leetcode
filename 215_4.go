package main

// 找数组的第k大的数据
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

func median3(a []int, i, j, k int) int {
	if a[i] < a[j] {
		if a[j] < a[k] {
			return j
		} else {
			if a[i] < a[k] {
				return k
			} else {
				return i
			}
		}
	} else { // a[i] > a[j]
		if a[j] > a[k] { // k,j,i
			return j
		} else {
			if a[i] < a[k] { // j, i,  k
				return i
			} else { // j,k, i,
				return k
			}
		}
	}
}

func partition(a []int, lo, hi int) int {
	// find median
	n := median3(a, lo, lo+(hi-lo+1)/2, hi)
	a[lo], a[n] = a[n], a[lo]

	//
	i, j := lo, hi+1
	v := a[lo]

	// a[lo] is max
	for i = i + 1; a[i] < v; i++ {
		if i == hi {
			a[lo], a[hi] = a[hi], a[lo]
			return hi
		}
	}

	// a[lo] is smallest
	for j = j - 1; v < a[j]; j-- {
		if j == lo+1 {
			return lo
		}
	}

	for i < j {
		a[i], a[j] = a[j], a[i]
		for {
			i++
			if a[i] >= v { // 大于等于时终止， 触发swap
				break
			}
		}
		for {
			j--
			if v >= a[j] {
				break
			}
		}
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func main() {
	var a = []int{3, 2, 1, 5, 6, 4}
	println(findKthLargest(a, 2))
}
