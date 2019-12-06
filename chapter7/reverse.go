package chapter7

func Reverse(node *ListNode) *ListNode {
	var previous *ListNode
	for node != nil {
		previous, node, node.Next = node, node.Next, previous
	}

	return previous
}
