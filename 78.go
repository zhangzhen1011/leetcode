// 返回所有子集
// 1
func subsets(nums []int) [][]int {
	var result = &[][]int{}
	var path = &[]int{}
	helper(nums, 0, path, result)
	return *result

}

func helper(nums []int, i int, path *[]int, result *[][]int) {
	tmp := make([]int, len(*path))
	copy(tmp, *path)
	*result = append(*result, tmp)
	if i > len(nums) {
		return
	}
	for s := i; s < len(nums); s++ {
		*path = append(*path, nums[s])
		helper(nums, s+1, path, result)
		*path = (*path)[:len(*path)-1]
	}
}

//2
// 返回所有子集, 区别是path
func subsets(nums []int) [][]int {
	var result = &[][]int{}
	helper(nums, 0, []int{}, result)
	return *result

}

func helper(nums []int, i int, path []int, result *[][]int) {
	tmp := make([]int, len(path))
	copy(tmp, path)
	*result = append(*result, tmp)
	if i >= len(nums) {
		return
	}
	for s := i; s < len(nums); s++ {
		path = append(path, nums[s])
		helper(nums, s+1, path, result)
		path = path[:len(path)-1]
	}
}

// 返回所有子集 , NOTICE 错误解法
func subsets(nums []int) [][]int {
	var result = &[][]int{}
	helper(nums, 0, []int{}, result)
	return *result

}

func helper(nums []int, i int, path []int, result *[][]int) {
	*result = append(*result, path) // 这里append时，所有递归的path slice底层都是同一个数组指针, 回溯过程会被覆盖，要用copy
	if i >= len(nums) {
		return
	}
	for s := i; s < len(nums); s++ {
		path = append(path, nums[s])
		helper(nums, s+1, path, result)
		path = path[:len(path)-1]
	}
}

// 方法3
func subsets(nums []int) [][]int {
	var result = [][]int{}
	dfs(nums, 0, []int{}, &result, true)
	return result
}

func dfs(nums []int, idx int, path []int, result *[][]int, diff bool) {
	if diff {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
	}
	if idx == len(nums) {
		return
	}

	// 选择idx位置
	dfs(nums, idx+1, append(path, nums[idx]), result, true)
	// 不选择idx位置
	dfs(nums, idx+1, path, result, false)
}
