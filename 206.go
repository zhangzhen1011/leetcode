/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 反转链表
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var h, p *ListNode
	for head != nil {
		p = head.Next
		head.Next = h
		h = head
		head = p
	}
	return h
}
