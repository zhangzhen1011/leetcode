/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	var m = map[*ListNode]bool{}
	for head != nil {
		_, ok := m[head]
		if ok {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}
