package chapter7

import (
	"testing"
)

var seven = ListNode{7, nil}
var five = ListNode{5, &seven}
var three = ListNode{3, &five}
var two = ListNode{2, &three}

func TestSearchListHit(t *testing.T) {
	result := SearchList(&two, 5)

	if result != &five {
		t.Fail()
	}
}

func TestSearchListMiss(t *testing.T) {
	result := SearchList(&two, 11)

	if result != nil {
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
