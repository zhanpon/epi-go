package chapter7

func Merge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head *ListNode
	if l1.Data < l2.Data {
		head = l1
		head.Next = Merge(l1.Next, l2)
	} else {
		head = l2
		head.Next = Merge(l1, l2.Next)
	}

	return head
}
