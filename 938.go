/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 遍历
func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	var sum = 0
	q := []*TreeNode{root}
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n.Val >= low && n.Val <= high {
			sum += n.Val
		}
		if n.Left != nil {
			q = append(q, n.Left)
		}
		if n.Right != nil {
			q = append(q, n.Right)
		}
	}

	return sum
}

