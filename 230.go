/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// bst的中根序是排好序的
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	result := &[]int{}
	inorder(root, result, k)
	return (*result)[k]
}

func inorder(node *TreeNode, result *[]int, k int) {
	if node == nil {
		return
	}
	if len(*result) > k {
		return
	}
	inorder(node.Left, result, k)
	*result = append(*result, node.Val)
	inorder(node.Right, result, k)
}
