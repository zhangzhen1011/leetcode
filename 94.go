/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 通用的方式解决遍历问题
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result, stack = make([]int, 0), make([]*opNode, 0)
	stack = append(stack, NewOpNode(0, root))
	for len(stack) > 0 {
		op := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if op != nil && op.Node != nil {
			if op.Op == 1 {
				result = append(result, op.Node.Val)
			} else {
				stack = append(stack, NewOpNode(0, op.Node.Right))
				stack = append(stack, NewOpNode(1, op.Node))
				stack = append(stack, NewOpNode(0, op.Node.Left))
			}
		}
	}

	return result
}

type opNode struct {
	Op   int // 0:visit or 1:print
	Node *TreeNode
}

func NewOpNode(op int, n *TreeNode) *opNode {
	return &opNode{op, n}
}
