package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 奇偶链表
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	list := make([]*ListNode, 0)
	for p := head; p != nil; p = p.Next {
		list = append(list, p)
	}
	var oH *ListNode
	var oT *ListNode
	var eH *ListNode
	var eT *ListNode
	for idx, item := range list {
		if idx%2 == 0 {
			if oH == nil {
				oH = item
				oT = item
			} else {
				oT.Next = item
				oT = item
			}
		}
		if idx%2 == 1 {
			if eH == nil {
				eH = item
				eT = item
			} else {
				eT.Next = item
				eT = item
			}
		}
	}
	oT.Next = eH
	eT.Next = nil
	return oH
}

func oddEvenListV2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	//3个或3个以上
	var oh, ot, eh, et *ListNode
	oh = head
	ot = head.Next.Next
	eh = head.Next
	et = eh
	oh.Next = ot
	et.Next = nil
	var idx = 0
	p := ot.Next
	ot.Next = nil
	for p != nil {
		if idx%2 == 0 {
			et.Next = p
			et = p
			p = p.Next
			et.Next = nil
		} else {
			ot.Next = p
			ot = p
			p = p.Next
			ot.Next = nil
		}
		idx++
	}
	ot.Next = eh
	return oh
}

// 头尾哨兵节点, 整体简洁
func oddEvenListV3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	var oh = &ListNode{}
	var eh = &ListNode{}
	var ot = oh
	var et = eh
	var isEven bool
	for head != nil {
		if isEven {
			et.Next = head
			et = et.Next
		} else {
			ot.Next = head
			ot = ot.Next
		}
		head = head.Next
		isEven = !isEven
	}
	ot.Next = eh.Next
	et.Next = nil
	return oh.Next
}

func main() {
	//h := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	h := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	ret := oddEvenListV3(h)
	for p := ret; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
