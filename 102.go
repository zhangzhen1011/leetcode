/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 按层输出
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	level := -1
	q := []*wrapNode{&wrapNode{0, root}}
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node.Level != level {
			ret = append(ret, []int{node.Val})
			level = node.Level
		} else {
			ret[len(ret)-1] = append(ret[len(ret)-1], node.Val)
		}

		if node.Left != nil {
			q = append(q, &wrapNode{node.Level + 1, node.Left})
		}
		if node.Right != nil {
			q = append(q, &wrapNode{node.Level + 1, node.Right})
		}
	}
	return ret
}

type wrapNode struct {
	Level int
	*TreeNode
}
