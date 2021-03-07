/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
// 复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	var m = map[*Node]*Node{}

	var newHead *Node = &Node{}
	var tail *Node = newHead
	for p := head; p != nil; p = p.Next {
		tail.Next = &Node{p.Val, p.Next, nil}
		tail = tail.Next
		m[p] = tail
	}
	p := head
	p2 := newHead.Next
	for p != nil {
		if p.Random == nil {
			p2.Random = nil
		} else {
			p2.Random = m[p.Random]
		}
		p = p.Next
		p2 = p2.Next
	}
	return newHead.Next
}

/////////////// 递归, 100%
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	if val, ok := m[head]; ok {
		return val
	}

	newH := &Node{head.Val, nil, nil}
	m[head] = newH
	newH.Next = copyRandomList(head.Next)
	newH.Random = copyRandomList(head.Random)
	return newH
}

var m = map[*Node]*Node{}
