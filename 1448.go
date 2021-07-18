/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
	count := 0
	helper(root, &count, -1e10)
	return count
}

func helper(node *TreeNode, count *int, maxVal int) {
	if node == nil {
		return
	}
	if node.Val >= maxVal {
		*count = *count + 1
		maxVal = node.Val
	}
	helper(node.Left, count, maxVal)
	helper(node.Right, count, maxVal)
}

//func max(i, j int) int {
//	if i > j {
//		return i
//	}
//	return j
//}
