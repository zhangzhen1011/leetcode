/**
 * Definition for singly-linked list.
  * type ListNode struct {
	   *     Val int
	    *     Next *ListNode
		 * }
*/
// leetcode 2 的变形,  考虑辅助空间，栈 ; 链表 要头插法
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var c = 0
	var s1 = make([]int, 0)
	var s2 = make([]int, 0)
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	var h = &ListNode{}
	p := h
	for len(s1) > 0 && len(s2) > 0 {
		val := (c + s1[len(s1)-1] + s2[len(s2)-1]) % 10
		c = (c + s1[len(s1)-1] + s2[len(s2)-1]) / 10
		p = h.Next
		h.Next = &ListNode{val, nil}
		h.Next.Next = p
		s1 = s1[:len(s1)-1]
		s2 = s2[:len(s2)-1]
	}
	for len(s1) > 0 {
		val := (c + s1[len(s1)-1]) % 10 // 这里可以再改改
		c = (c + s1[len(s1)-1]) / 10
		p = h.Next
		h.Next = &ListNode{val, nil}
		h.Next.Next = p
		s1 = s1[:len(s1)-1]
	}
	for len(s2) > 0 {
		val := (c + s2[len(s2)-1]) % 10
		c = (c + s2[len(s2)-1]) / 10
		p = h.Next
		h.Next = &ListNode{val, nil}
		h.Next.Next = p
		s2 = s2[:len(s2)-1]
	}
	if c == 1 {
		p = h.Next
		h.Next = &ListNode{1, nil}
		h.Next.Next = p
	}
	return h.Next
}
