package chapter7

func HasCycle(head *ListNode) *ListNode {
	count := make(map[*ListNode]int)

	for node := head; node != nil; node = node.Next {
		if count[node] != 0 {
			return node
		}

		count[node]++
	}

	return nil
}
