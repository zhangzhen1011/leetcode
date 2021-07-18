package main

import (
	"fmt"
	"strconv"
	"strings"
)

// * Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	var ret string = strconv.Itoa(root.Val)
	ret += ","
	ret += this.serialize(root.Left)
	ret += ","
	ret += this.serialize(root.Right)
	return ret
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	idx := 0
	ss := strings.Split(data, ",")
	return helper(ss, &idx)
}

func helper(data []string, idx *int) *TreeNode {
	if *idx >= len(data) {
		return nil
	}
	if data[*idx] == "#" {
		*idx = *idx + 1
		return nil
	}

	val, _ := strconv.Atoi(data[*idx])
	node := &TreeNode{val, nil, nil}
	*idx = *idx + 1
	node.Left = helper(data, idx)
	node.Right = helper(data, idx)
	return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
func main() {
	c := Constructor()
	r := &TreeNode{11, &TreeNode{2, nil, nil}, &TreeNode{3, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}}
	//r := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	ret := c.serialize(r)
	fmt.Println(ret)
	fmt.Println(c.deserialize(ret))
}
