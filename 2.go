/**
 * Definition for singly-linked list.
  * type ListNode struct {
	   *     Val int
	    *     Next *ListNode
		 * }
*/
// 两链表相加，逆序
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c = 0
	var h = &ListNode{}
	var p = h
	for l1 != nil && l2 != nil {
		val := (c + l1.Val + l2.Val) % 10
		c = (l1.Val + l2.Val + c) / 10
		p.Next = &ListNode{val, nil}
		p = p.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l2 != nil {
		val := (c + l2.Val) % 10
		c = (l2.Val + c) / 10
		p.Next = &ListNode{val, nil}
		p = p.Next
		l2 = l2.Next
	}
	for l1 != nil {
		val := (c + l1.Val) % 10
		c = (l1.Val + c) / 10
		p.Next = &ListNode{val, nil}
		p = p.Next
		l1 = l1.Next
	}
	if c == 1 {
		p.Next = &ListNode{1, nil}
	}
	return h.Next
}
