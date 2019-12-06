package chapter7

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var (
		three = ListNode{3, nil}
		two   = ListNode{2, &three}
		one   = ListNode{1, &two}
	)

	result := Reverse(&one)

	if result != &three || result.Next != &two || result.Next.Next != &one {
		t.Fail()
	}
}
