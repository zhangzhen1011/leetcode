/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 最近公共祖先
// 1. 遍历， map记录路径
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	var m = map[*TreeNode]*TreeNode{}
	var vis = map[*TreeNode]bool{} // 是否访问标记
	m[root] = nil

	// bfs
	q := []*TreeNode{root}
	for len(q) > 0 {
		h := q[0]
		n = q[1:]
		vis[n] = false
		if n.Left != nil {
			m[n.Left] = n
			q = append(q, n.Left)
		}
		if n.Right != nil {
			m[n.Right] = n
			q = append(q, n.Right)
		}
	}

	vis[p] = true
	for {
		val := m[p]
		if val == nil {
			break
		}
		vis[val] = true
		p = val
	}
	vis[q] = true
	for {
		val := m[q]
		if val == nil {
			break
		}
		if vis[val] {
			return val
		}
		p = val
	}
	return nil
}

// 2. 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left == nil && right == nil {
		return root
	}
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return root
}
