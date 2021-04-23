// n叉树和二叉树的转换（decode，encode）
// 多叉树 -> 二叉树: 保留大儿子，大儿子串联所有兄弟
// 二叉树 -> 多叉树: 逆向运算

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) encode(root *Node) *TreeNode {
	if root == nil {
		return root
	}
	node := &TreeNode{Val: root.Val}
	if len(root.Children) == 0 {
		return node
	}

	node.Left = this.encode(root.Children[0])
	h := node.Left
	for i := 1; i < len(root.Children); i++ {
		h.Right = this.encode(root.Children[i])
		h = h.Right
	}
	return node
}

func (this *Codec) decode(root *TreeNode) *Node {
	if root == nil {
		return nil
	}

	node := &Node{Val: root.Val}
	if root.Left == nil {
		return node
	}

	//node.Children = append(node.Children, this.decode(root.Left))
	h := root.Left
	for h != nil {
		node.Children = append(node.Children, this.decode(h))
		h = h.Right
	}

	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * bst := obj.encode(root);
 * ans := obj.decode(bst);
 */
