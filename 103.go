/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var result, queue = make([][]int, 0), make([]*wrapNode, 0)
	queue = append(queue, &wrapNode{0, root})
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		if n.Node != nil {
			if len(result) < n.Level+1 {
				result = append(result, []int{})
			}

			if n.Level%2 == 0 {
				result[n.Level] = append(result[n.Level], n.Node.Val)
			} else {
				result[n.Level] = append([]int{n.Node.Val}, result[n.Level]...)
			}
			queue = append(queue, &wrapNode{n.Level + 1, n.Node.Left})
			queue = append(queue, &wrapNode{n.Level + 1, n.Node.Right})
		}
	}

	return result
}

type wrapNode struct {
	Level int
	Node  *TreeNode
}

