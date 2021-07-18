package main

// 两个正序数组的中位数

import "math"

// 两个正序数组的中位数
// 二分
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	if n1 > n2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	lo, hi := 0, n1
	var p1, p2 int
	for lo <= hi {
		p1 = (lo + hi) / 2
		p2 = (n1+n2+1)/2 - p1

		var maxLeft1, minRight1, maxLeft2, minRight2 int
		if p1 <= 0 {
			maxLeft1 = math.MinInt64
		} else {
			maxLeft1 = nums1[p1-1]
		}
		if p1 >= n1 {
			minRight1 = math.MaxInt64
		} else {
			minRight1 = nums1[p1]
		}
		if p2 <= 0 {
			maxLeft2 = math.MinInt64
		} else {
			maxLeft2 = nums2[p2-1]
		}
		if p2 >= n2 {
			minRight2 = math.MaxInt64
		} else {
			minRight2 = nums2[p2]
		}

		if maxLeft1 <= minRight2 && maxLeft2 <= minRight1 {
			if (n1+n2)%2 == 0 {
				return float64(max(maxLeft1, maxLeft2)+min(minRight1, minRight2)) / float64(2)
			} else {
				return float64(max(maxLeft1, maxLeft2))
			}
		} else if maxLeft1 > minRight2 {
			hi = p1 - 1
		} else {
			lo = p1 + 1
		}
	}
	panic("not found")
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 第二种解法，一次遍历
// 双指针 按顺序遍历
func findMedianSortedArrays(nums1, nums2 []int) float64 {
	n1 := len(nums1)
	n2 := len(nums2)
	var i, j = 0, 0
	var cur int = -1e10
	var prev int = -1e10
	for e := 0; e <= (n1+n2)/2; e++ {
		prev = cur
		if (i < n1 && j >= n2) || (i < n1 && nums1[i] < nums2[j]) {
			cur = nums1[i]
			i++
		} else {
			cur = nums2[j]
			j++
		}
	}
	if (n1+n2)%2 == 0 {
		return float64(prev+cur) / 2.0
	} else {
		return float64(cur)
	}
}
