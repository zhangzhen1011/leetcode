package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	l1 := 0
	l2 := 0
	for p := headA; p != nil; p = p.Next {
		l1++
	}
	for p := headB; p != nil; p = p.Next {
		l2++
	}
	for l1 > l2 {
		headA = headA.Next
		l1--
	}
	for l1 < l2 {
		headB = headB.Next
		l2--
	}
	for headA != headB {
		headA = headA.Next
		headB = headB.Next
	}
	return headB
}

func main() {

	l5 := &ListNode{
		5, nil,
	}
	l4 := &ListNode{
		4, l5,
	}

	l8 := &ListNode{
		8, l4,
	}
	h1 := &ListNode{
		4, &ListNode{
			1, l8,
		},
	}
	h2 := &ListNode{
		5, &ListNode{
			6, &ListNode{
				1, l8,
			},
		},
	}

	ret := getIntersectionNode(h1, h2)
	fmt.Println(ret.Val)
}

type ListNode struct {
	Val  int
	Next *ListNode
}
