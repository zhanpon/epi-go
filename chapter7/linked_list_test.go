package chapter7

import (
	"testing"
)

func TestSearchListHit(t *testing.T) {
	var (
		seven = ListNode{7, nil}
		five  = ListNode{5, &seven}
		three = ListNode{3, &five}
		two   = ListNode{2, &three}
	)

	if SearchList(&two, 5) != &five {
		t.Fail()
	}

	if SearchList(&two, 11) != nil {
		t.Fail()
	}
}

func TestInsertAfter(t *testing.T) {
	var (
		three = ListNode{3, nil}
		one   = ListNode{1, &three}
	)

	two := ListNode{2, nil}
	InsertAfter(&one, &two)

	if one.Next != &two {
		t.Fail()
	}
}
