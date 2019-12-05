package chapter7

import "testing"

var (
	seven = ListNode{7, nil}
	five  = ListNode{5, &seven}
	two   = ListNode{2, &five}
)

var (
	eleven = ListNode{11, nil}
	three  = ListNode{3, &eleven}
)

func TestMerge(t *testing.T) {
	node := Merge(&two, &three)

	expected := []int{2, 3, 5, 7, 11}

	for _, i := range expected {
		if i != node.Data {
			t.Fatalf("expected %v, got %v", i, node.Data)
		}

		node = node.Next
	}
}
