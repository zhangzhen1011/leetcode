/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	q := []*wrapNode{&wrapNode{root, 0}}

	//stack := []*Node{}

	level := -1
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if len(q) == 0 {
			node.Next = nil
		} else {
			if node.Level == q[0].Level {
				node.Next = q[0].Node
			} else {
				node.Next = nil
			}
		}
		if level != node.Level {
			//stack = append(stack, nil)
			level = node.Level
		}
		//stack = append(stack, node.Node)
		if node.Left != nil {
			q = append(q, &wrapNode{node.Left, node.Level + 1})
		}
		if node.Right != nil {
			q = append(q, &wrapNode{node.Right, node.Level + 1})
		}
	}

	//for i := len(stack) - 1; i > 0; i-- {
	//	if stack[i-1] != nil {
	//		stack[i-1].Next = stack[i]
	//	}
	//}
	return root
}

type wrapNode struct {
	*Node
	Level int
}
