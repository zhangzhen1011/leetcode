/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// bst 转化成递增顺序搜索树
// 中序遍历时同时构建
func increasingBST(root *TreeNode) *TreeNode {
	inorder(root, nil)
	return r
}

var r *TreeNode

func inorder(node *TreeNode, node2 *TreeNode) *TreeNode {
	if node == nil {
		return node2
	}
	node2 = inorder(node.Left, node2)
	if node2 == nil {
		node2 = &TreeNode{Val: node.Val}
		r = node2
	} else {
		node2.Right = &TreeNode{Val: node.Val}
		node2 = node2.Right
	}
	return inorder(node.Right, node2)
}
