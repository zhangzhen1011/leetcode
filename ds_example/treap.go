package main

import (
	"fmt"
	"math/rand"
)

type node struct {
	ch       [2]*node // 0,左，1右
	priority int      // 堆优先级, 保持平衡
	val      int
}

func (o *node) cmp(b int) int {
	switch {
	case b < o.val: // 插入左子树
		return 0
	case b > o.val: // 插入右子树
		return 1
	default: // 不操作
		return -1
	}
}

// 旋转, 左旋右旋合并, d^1  ,1-d
// 假设d=0
func (o *node) rotate(d int) *node {
	x := o.ch[d^1]      // d=0, 左旋, d^1=1, 取出右孩子
	o.ch[d^1] = x.ch[d] // 当前节点的右孩子 = 前右孩子的左孩子
	x.ch[d] = o         // 前右孩子的左孩子 = 当前节点
	return x            // 返回左旋有的节点
}

// NOTICE假设无重复节点
type treap struct {
	root *node
}

// 插入, 保证bst的顺序性，同时保证优先级的堆性质
// 向当前子树插入，并返回子树根节点
func (t *treap) _put(o *node, val int) *node {
	if o == nil { // 递归出口
		return &node{priority: rand.Int(), val: val}
	}

	if d := o.cmp(val); d >= 0 {
		o.ch[d] = t._put(o.ch[d], val)
		// 是否旋转
		if o.ch[d].priority > o.priority {
			o = o.rotate(d ^ 1) // 插入左子树，则右旋; 插入右子树，则左旋
		}
	}
	return o
}

func (t *treap) put(val int) {
	t.root = t._put(t.root, val)
}

// 后驱, 比当前值大的最近的node
func (t *treap) lowerBound(val int) (lb *node) {
	for o := t.root; o != nil; {
		switch c := o.cmp(val); {
		case c == 0:
			lb = o
			o = o.ch[0]
		case c > 0:
			o = o.ch[1]
		default:
			return o
		}
	}
	return
}

func (t *treap) inorder(o *node) {
	//var ret []int
	if o == nil {
		return
	}
	t.inorder(o.ch[0])
	fmt.Println(o.val)
	t.inorder(o.ch[1])
}

func main() {
	t := &treap{}
	t.put(2)
	t.put(0)
	t.put(4)
	t.put(-2)
	t.put(6)
	//t.inorder(t.root)
	fmt.Println(t.lowerBound(3).val)
}
