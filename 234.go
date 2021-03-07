
// 判断回文链表, 栈
// 耗时比较多
func isPalindrome(head *ListNode) bool {
	stack := make([]*ListNode, 0)
	p := head
	for p != nil {
		stack = append(stack, p)
		p = p.Next
	}
	for i := len(stack) - 1; i >= 0 && head != nil; i-- {
		if stack[i].Val != head.Val {
			return false
		}
		head = head.Next
	}
	return true
}

// TODO, 减少压栈
func isPalindromeV2(head *ListNode) bool {
}

// TODO, 不使用栈， 结合链表反转
func isPalindromeV3(head *ListNode) bool {
}
