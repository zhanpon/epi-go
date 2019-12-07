package chapter7

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	var (
		three = ListNode{3, nil}
		two   = ListNode{2, &three}
		one   = ListNode{1, &two}
	)

	result := HasCycle(&one)

	if result != nil {
		t.Fail()
	}
}

func TestHasCycle2(t *testing.T) {
	var (
		four  = ListNode{4, nil}
		three = ListNode{3, &four}
		two   = ListNode{2, &three}
		one   = ListNode{1, &two}
	)
	four.Next = &two

	result := HasCycle(&one)

	if result != &two {
		t.Fail()
	}
}
