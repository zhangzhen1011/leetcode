/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var c = 0
	p := head
	prev := &ListNode{}
	prevprev := &ListNode{}
	prev.Next = head
	prevprev.Next = prev
	h := prev
	for p != nil {
		if c%2 == 1 {
			//next := p.Next
			prev.Next = p.Next
			p.Next = prev
			prevprev.Next = p
			p, prev = prev, p
		}
		p = p.Next
		prev = prev.Next
		prevprev = prevprev.Next
		c++
	}
	return h.Next
}
