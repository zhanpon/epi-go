package chapter7

func Reverse(node *ListNode) *ListNode {
	var last *ListNode
	for node != nil {
		next := node.Next
		node.Next = last
		last = node
		node = next
	}

	return last
}
