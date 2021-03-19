package main

import "fmt"

func permute(nums []int) [][]int {
	var result = &[][]int{}
	var visited = make([]bool, len(nums))
	helper(nums, &visited, []int{}, result)
	return *result
}

func helper(nums []int, visited *[]bool, path []int, result *[][]int) {
	if len(nums) == len(path) {
		*result = append(*result, path)
		return
	}
	for i := 0; i < len(nums); i++ {
		if !(*visited)[i] {
			(*visited)[i] = true
			helper(nums, visited, append(path, nums[i]), result)
			(*visited)[i] = false
		}
	}
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(permute(a))
}
