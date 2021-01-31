package main

import "fmt"

type ListNode struct {
	Key  int
	Val  int
	Prev *ListNode
	Next *ListNode
}

type LRUCache struct {
	m  map[int]*ListNode
	dl *duList

	capacity int
}

func Constructor(capacity int) LRUCache {
	ret := LRUCache{
		m:        make(map[int]*ListNode),
		dl:       initDuList(),
		capacity: capacity,
	}

	return ret
}

// -1, not found
func (this *LRUCache) Get(key int) int {
	node, ok := this.m[key]
	if ok {
		this.dl.mvHead(node)
		return node.Val
	}
	return -1
}

// always keep key first
func (this *LRUCache) Put(key int, value int) {
	if old, ok := this.m[key]; ok {
		// mod old, mv to first
		old.Val = value
		this.dl.mvHead(old)
	} else {
		if len(this.m) >= this.capacity { // full
			// del most old
			last := this.dl.removeLast()
			delete(this.m, last.Key)
		}
		// add new
		newVal := &ListNode{
			Key: key,
			Val: value,
		}
		this.dl.addFirst(newVal)
		this.m[key] = newVal
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type duList struct {
	dummyHead *ListNode
	dummyTail *ListNode
}

func initDuList() *duList {
	d := &duList{
		dummyHead: &ListNode{},
		dummyTail: &ListNode{},
	}
	d.dummyHead.Next = d.dummyTail
	d.dummyTail.Prev = d.dummyHead
	return d
}

func (d *duList) addFirst(x *ListNode) {
	x.Prev = d.dummyHead
	x.Next = d.dummyHead.Next
	d.dummyHead.Next.Prev = x // 漏了
	d.dummyHead.Next = x
}

func (d *duList) remove(x *ListNode) {
	x.Prev.Next = x.Next
	x.Next.Prev = x.Prev
}

func (d *duList) removeLast() *ListNode {
	tail := d.dummyTail.Prev
	if tail != d.dummyHead {
		d.remove(tail)
		return tail
	}
	return nil
}

func (d *duList) mvHead(x *ListNode) {
	d.remove(x)
	d.addFirst(x)
}

func (d *duList) printAll() {
	for s := d.dummyHead.Next; s != d.dummyTail; s = s.Next {
		println("--", s.Key, s.Val)
	}
}

func main() {
	cache := Constructor(2)
	cache.Put(2, 1)
	cache.Put(2, 2)
	cache.dl.printAll()
	fmt.Println(cache.Get(2))
	cache.dl.printAll()
	cache.Put(1, 1)
	//fmt.Println(cache.Get(2))
	cache.Put(4, 1)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(2))
	cache.dl.printAll()
}
