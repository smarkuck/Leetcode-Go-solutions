package leetcode

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	head := new(ListNode)
	cur := head
	for l1 != nil || l2 != nil || carry != 0 {
		n1, n2 := 0, 0
		if l1 != nil {
			n1, l1 = l1.Val, l1.Next
		}
		if l2 != nil {
			n2, l2 = l2.Val, l2.Next
		}
		num := n1 + n2 + carry
		carry = num / 10
		cur.Next = &ListNode{Val: num % 10, Next: nil}
		cur = cur.Next
	}
	return head.Next
}

func toArray(l *ListNode) []int {
	a := []int{}
	for l != nil {
		a = append(a, l.Val)
		l = l.Next
	}
	return a
}

func TestAddTwoNumbers(t *testing.T) {
	input1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	input2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	output := addTwoNumbers(input1, input2)
	expected := &ListNode{7, &ListNode{0, &ListNode{8, nil}}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toArray(output), toArray(expected))
	}

	input1 = &ListNode{0, nil}
	input2 = &ListNode{0, nil}
	output = addTwoNumbers(input1, input2)
	expected = &ListNode{0, nil}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toArray(output), toArray(expected))
	}

	input1 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9,
		&ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}}
	input2 = &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}
	output = addTwoNumbers(input1, input2)
	expected = &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9,
		&ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toArray(output), toArray(expected))
	}
}
