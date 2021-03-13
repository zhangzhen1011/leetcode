/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 前序+后序，重建树
func buildTree(preorder []int, inorder []int) *TreeNode {
	var inPos = map[int]int{}
	for idx, v := range inorder {
		inPos[v] = idx
	}
	return buildDFS(preorder, 0, inPos)
}

func buildDFS(pre []int, inStart int, inPos map[int]int) *TreeNode {
	if len(pre) == 0 {
		return nil
	}
	root := &TreeNode{pre[0], nil, nil}
	rootIdx := inPos[pre[0]]
	leftLen := rootIdx - inStart
	root.Left = buildDFS(pre[1:leftLen+1], inStart, inPos)
	root.Right = buildDFS(pre[leftLen+1:], rootIdx+1, inPos)
	return root
}
