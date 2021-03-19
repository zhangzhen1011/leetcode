// 返回所有子集
func subsets(nums []int) [][]int {
	var result = &[][]int{}
	var path = *[]int{} // 这个地方要用指针。  因为下探过程中, path底层数组会grow
	helper(nums, 0, []int{}, result)
	return *result

}

func helper(nums []int, i int, path *[]int, result *[][]int) {
	tmp := make([]int, len(*path))
	copy(tmp, *path)
	*result = append(*result, tmp)
	if i >= len(nums) {
		return
	}
	for s := i + 1; s < len(nums); s++ {
		*path = append(*path, nums[i])
		helper(nums, s, path, result)
		*path = *path[:len(*path)-1]
	}
}
