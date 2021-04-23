package main

import (
	"container/heap"
	"fmt"
	"math/rand"
)

// 矩阵中找子矩形
// 子矩阵：通过两个点定位（左上角，右下角）
// bf: 暴力搜索所有可能, 4个for，遍历边界，O(m*m*n*n)
// 1. 矩阵在某个维度拍扁，转化为一维的前缀和
// 2. 高效排序，优先队列
// 3. 一维数组的子数组和的问题（前缀数组解决）
func maxSumSubmatrix(matrix [][]int, k int) int {
	m := len(matrix)
	n := len(matrix[0])
	//fmt.Println(m, n)
	if m > n { // 减少外层，内部logN
		m2 := make([][]int, n)
		for i := 0; i < len(m2); i++ {
			m2[i] = make([]int, m)
		}

		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				m2[j][i] = matrix[i][j]
			}
		}

		return maxSumSubmatrix(m2, k)
	}

	var ret int = -1e10
	// 固定子区域上下边界
	for i := 0; i < m; i++ {
		row := make([]int, n)
		for j := i; j < m; j++ {
			for _k := 0; _k < n; _k++ {
				row[_k] += matrix[j][_k] // 每列和累加
			}
			ret = max(ret, helper(row, k))
		}
	}
	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// treap
func helper(row []int, k int) int {
	t := &treap{}
	t.put(0) // NOTE 这不不能少 ,  因为存储的前缀和，最初的前缀和可以理解为是0
	var presum int
	var ret int = -1e10
	for i := 0; i < len(row); i++ {
		presum += row[i]
		o := t.lowerBound(presum - k)
		if o != nil {
			ret = max(ret, presum-o.val)
		}
		t.put(presum)
	}
	return ret
}

// 优先队列耗时多，超时了
func helper(row []int, k int) int {
	//fmt.Println(row)
	var p pq // 存储pre[i]
	heap.Push(&p, 0)
	var presum int
	var ret int = -1e10
	for i := 0; i < len(row); i++ {
		presum += row[i]
		var tmp int = -1e10
		var tmpPre pq = make([]int, len(p))
		copy(tmpPre, p)
		for len(tmpPre) > 0 {
			x := heap.Pop(&tmpPre).(int)
			if presum-x <= k {
				tmp = x
				break
			}
		}

		if tmp != -1e10 {
			ret = max(ret, presum-tmp)
		}

		heap.Push(&p, presum) // pre[j] - pre[i]<= k;  pre[i]>=pre[j] -k
	}
	return ret
}

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

type pq []int

func (p pq) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p *pq) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *pq) Len() int {
	return len(*p)
}

func (p *pq) Push(x interface{}) {
	xx := x.(int)
	*p = append(*p, xx)
}

func (p *pq) Pop() interface{} {
	old := *p
	*p = old[:len(old)-1]
	return old[len(old)-1]
}

func main() {
	//var p pq
	//heap.Push(&p, 1)
	//heap.Push(&p, -1)
	//heap.Push(&p, 2)
	//fmt.Println(p)
	//heap.Pop(&p)
	//heap.Push(&p, 3)
	//fmt.Println(p)
	//heap.Push(&p, -2)
	//fmt.Println(p)

	fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
}
