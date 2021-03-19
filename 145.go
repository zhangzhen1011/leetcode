/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//wrap节点
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result, stack = make([]int, 0), make([]*wrapNode, 0)
	stack = append(stack, &wrapNode{0, root})
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if n.Node != nil {
			if n.Op == 1 {
				result = append(result, n.Node.Val)
			} else {
				stack = append(stack, &wrapNode{1, n.Node})
				stack = append(stack, &wrapNode{0, n.Node.Right})
				stack = append(stack, &wrapNode{0, n.Node.Left})
			}
		}
	}
	return result
}

type wrapNode struct {
	Op   int // 0, visit , 1 print
	Node *TreeNode
}
