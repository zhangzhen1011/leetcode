/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDiffInBST(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var ret int = 1e13
	var path []int
	dfs(root, &path, &ret)
	return ret
}

func dfs(node *TreeNode, nums *[]int, result *int) {
	if node == nil {
		return
	}

	for _, val := range *nums {
		*result = min(*result, abs(val-node.Val))
	}

	*nums = append(*nums, node.Val)
	dfs(node.Left, nums, result)
	dfs(node.Right, nums, result)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
