package main

import "fmt"

///func sortColors(nums []int) {
///	sort(nums, 0, len(nums)-1)
///}
///
///func partition(nums []int, lo, hi int) {
///	var i, j := lo+1, hi
///	for i < j {
///		for i <= hi {
///			if nums[i] > nums[lo] {
///				break
///			}
///			i++
///		}
///		for j > lo {
///			if nums[j] < nums[lo] {
///				break
///			}
///			j--
///		}
///
///		if i > j {
///			break
///		}
///		nums[i], nums[j] = nums[j], nums[i]
///	}
///	if j > lo {
///		nums[lo], nums[j] = nums[j], nums[lo]
///	}
///	return j
///}
///
///func sort(nums []int, lo, hi int) {
///	p := partition(nums, 0, len(nums)-1)
///	sort(nums, 0, p-1)
///	sort(nums, p+1, len(nums)-1)
///}

// 快排的另一个写法
func sortColors(nums []int) {
	sort(nums, 0, len(nums)-1)
}

func partition(nums []int, lo, hi int) int {
	i := lo - 1
	anchor := nums[hi]
	for j := lo; j < hi; j++ {
		if nums[j] <= anchor {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[hi], nums[i+1] = nums[i+1], nums[hi]
	return i + 1
}

func sort(nums []int, lo, hi int) {
	if lo < hi {
		p := partition(nums, 0, hi)
		sort(nums, 0, p-1)
		sort(nums, p+1, hi)
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
