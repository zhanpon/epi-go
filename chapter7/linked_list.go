package chapter7

type ListNode struct {
	Data int
	Next *ListNode
}

func SearchList(node *ListNode, key int) *ListNode {
	for node != nil && node.Data != key {
		node = node.Next
	}

	return node
}
