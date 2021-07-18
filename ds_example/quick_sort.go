package main

import "fmt"

func quickSort(nums []int, lo, hi int) {
	if lo < hi {
		j := partition(nums, lo, hi)
		fmt.Println(j)
		quickSort(nums, lo, j-1)
		quickSort(nums, j+1, hi)
	}
}

func partition(nums []int, lo, hi int) int {
	anchor := nums[hi] // 最后一个数为anchor
	i := lo - 1
	for j := lo; j <= hi-1; j++ { // 一次遍历, hi-1
		if nums[j] <= anchor {
			i++
			nums[i], nums[j] = nums[j], nums[i] // 比anchor小的挪到左边
		}
	}
	// j == hi, anchor归位
	nums[i+1], nums[hi] = nums[hi], nums[i+1] // i+1
	return i + 1
}

func main() {
	a := []int{2, 0, 2, 1, 1, 0}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
