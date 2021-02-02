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

func partition(a []int, lo, hi int) int {
	i, j := lo, hi+1
	p := a[lo]
	for {
		for i < hi {
			i = i + 1
			if a[i] < p {
				continue
			}
			break
		}
		for lo < j {
			j = j - 1
			if p < a[j] {
				continue
			}
			break
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func main() {
	var a = []int{3, 2, 1, 5, 6, 4}
	println(findKthLargest(a, 2))
}
