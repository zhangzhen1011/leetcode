/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	var ret []int
	var s = make([]*TreeNode, 0)
	s = append(s, root)
	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]
		for node != nil {
			if node.Right != nil {
				s = append(s, node.Right)
			}
			ret = append(ret, node.Val)
			node = node.Left
		}
	}
	return ret
}
