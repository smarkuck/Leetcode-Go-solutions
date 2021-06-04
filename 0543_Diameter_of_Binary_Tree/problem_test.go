package leetcode

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
	diameter = 0
	longestPath(root)
	return diameter
}

func longestPath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := longestPath(root.Left)
	right := longestPath(root.Right)
	pathSize := left + right
	if pathSize > diameter {
		diameter = pathSize
	}
	if left > right {
		return left + 1
	}
	return right + 1
}

func TestDiameterOfBinaryTree(t *testing.T) {
	tree := &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil}},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil}}

	output := diameterOfBinaryTree(tree)
	expected := 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	tree = &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: nil, Right: nil},
		Right: nil}

	output = diameterOfBinaryTree(tree)
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
