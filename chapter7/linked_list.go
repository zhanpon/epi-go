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

func InsertAfter(node *ListNode, newNode *ListNode) {
	node.Next, newNode.Next = newNode, node.Next
}

func DeleteList(node *ListNode) {
	node.Next = node.Next.Next
}
