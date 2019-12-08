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

// Less memory but quadratic time
func HasCycleQuadratic(head *ListNode) *ListNode {
	var n int
	for node1 := head; node1 != nil; node1 = node1.Next {
		node2 := head
		for i := 0; i < n; i++ {
			if node1 == node2 {
				return node1
			}
			node2 = node2.Next
		}
		n++
	}

	return nil
}
