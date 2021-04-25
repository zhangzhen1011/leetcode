// 搜索旋转数组
// 二分
package main

import "fmt"

func search(nums []int, target int) int {
	return search_(nums, 0, len(nums)-1, target)
}

func search_(nums []int, lo, hi int, target int) int {
	if lo <= hi {
		mid := (lo + hi) / 2
		if nums[mid] == target {
			return mid
		}
		if target >= nums[lo] {
			if nums[mid] >= nums[lo] {
				if target < nums[mid] {
					return search_(nums, lo, mid-1, target)
				} else {
					return search_(nums, mid+1, hi, target)
				}
			} else {
				return search_(nums, lo, mid-1, target)
			}
		} else { // 后半段
			if nums[mid] >= nums[lo] {
				return search_(nums, mid+1, hi, target)
			} else { // mid在后半段
				if target < nums[mid] {
					return search_(nums, lo, mid-1, target)
				} else {
					return search_(nums, mid+1, hi, target)
				}
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
