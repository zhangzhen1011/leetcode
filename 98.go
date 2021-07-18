/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 验证是否有效的二叉搜索树，  中根遍历，  看是否有序
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var prev int = -1e10
	return isValid(root, &prev)
}

// NOTICE  prev  要传指针，需要在整个递归期间同步
func isValid(node *TreeNode, prev *int) bool {
	if node == nil {
		return true
	}

	if !isValid(node.Left, prev) {
		return false
	}

	if node.Val <= *prev {
		return false
	}
	*prev = node.Val

	return isValid(node.Right, prev)
}
