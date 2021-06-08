package leetcode

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{stack: []*TreeNode{}}
	for root != nil {
		it.stack = append(it.stack, root)
		root = root.Left
	}
	return it
}

func (it *BSTIterator) Next() int {
	leftmostNode := it.stack[len(it.stack)-1]
	it.stack = it.stack[:len(it.stack)-1]
	node := leftmostNode.Right
	for node != nil {
		it.stack = append(it.stack, node)
		node = node.Left
	}
	return leftmostNode.Val
}

func (it *BSTIterator) HasNext() bool {
	return len(it.stack) > 0
}

func TestBSTIterator(t *testing.T) {
	it := Constructor(&TreeNode{Val: 7,
		Left: &TreeNode{Val: 3, Left: nil, Right: nil},
		Right: &TreeNode{Val: 15,
			Left:  &TreeNode{Val: 9, Left: nil, Right: nil},
			Right: &TreeNode{Val: 20, Left: nil, Right: nil}}})

	output := it.Next()
	expected := 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = it.Next()
	expected = 7
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output2 := it.HasNext()
	expected2 := true
	if output2 != expected2 {
		t.Errorf("\noutput:   %v\nexpected: %v", output2, expected2)
	}

	output = it.Next()
	expected = 9
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output2 = it.HasNext()
	expected2 = true
	if output2 != expected2 {
		t.Errorf("\noutput:   %v\nexpected: %v", output2, expected2)
	}

	output = it.Next()
	expected = 15
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output2 = it.HasNext()
	expected2 = true
	if output2 != expected2 {
		t.Errorf("\noutput:   %v\nexpected: %v", output2, expected2)
	}

	output = it.Next()
	expected = 20
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output2 = it.HasNext()
	expected2 = false
	if output2 != expected2 {
		t.Errorf("\noutput:   %v\nexpected: %v", output2, expected2)
	}
}
