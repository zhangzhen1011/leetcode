package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 设计跳表
type Skiplist struct {
	head *node
	//tail     *node
	maxLevel int // 记录跳表中所有节点的最大高度
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const Max_level = 16
const Max_value = 9e12

type node struct {
	val   int
	nexts []*node // 每层的后继节点
}

func randomLevel() int {
	l := 1
	for rand.Intn(2) == 1 && l < Max_level {
		l++
	}
	return l
}

func Constructor() Skiplist {
	s := Skiplist{
		head: &node{
			val:   -Max_value,
			nexts: make([]*node, Max_level), // 每一层的哨兵
		},
		maxLevel: 1,
	}
	return s
}

func (this *Skiplist) Search(target int) bool {
	prev := this.search(target)[0] // 获取第0层前驱节点
	return prev.nexts[0] != nil && prev.nexts[0].val == target
}

func (this *Skiplist) search(key int) []*node {
	cur := this.head
	prevs := make([]*node, Max_level)
	// 从每个节点的顶层开始
	for i := this.maxLevel - 1; i >= 0; i-- {
		// 满足条件则继续
		for cur.nexts[i] != nil && cur.nexts[i].val < key {
			cur = cur.nexts[i]
		}
		// 内层循环终止
		prevs[i] = cur // 存储每一层的前驱节点
	}
	return prevs
}

func (this *Skiplist) Add(num int) {
	prevs := this.search(num)
	levelNum := randomLevel() // 随机每个节点的高度
	if levelNum > this.maxLevel {
		for i := this.maxLevel; i < levelNum; i++ {
			prevs[i] = this.head
		}
		this.maxLevel = levelNum
	}

	cur := &node{
		val:   num,
		nexts: make([]*node, levelNum),
	}
	// 从上层开始, 插入节点，调整每一层的指针
	for i := levelNum - 1; i >= 0; i-- {
		cur.nexts[i] = prevs[i].nexts[i]
		prevs[i].nexts[i] = cur
	}
}

func (this *Skiplist) Erase(num int) bool {
	prevs := this.search(num)
	if prevs[0] == nil {
		return false
	}
	if prevs[0].nexts[0] == nil || prevs[0].nexts[0].val != num {
		return false
	}
	del := prevs[0].nexts[0]
	for i := 0; i < this.maxLevel; i++ {
		if prevs[i].nexts[i] == del {
			prevs[i].nexts[i] = del.nexts[i]
		}
	}
	// 更新maxlevel
	for this.maxLevel > 1 && this.head.nexts[this.maxLevel-1] == nil {
		this.maxLevel--
	}
	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

func main() {
	sl := Constructor()
	sl.Add(1)
	sl.Add(2)
	sl.Add(4)
	fmt.Println(sl.Search(0))
	sl.Add(3)
	fmt.Println(sl.Search(1))
	sl.Erase(1)
	fmt.Println(sl.Search(1))
	sl.Add(0)
}
