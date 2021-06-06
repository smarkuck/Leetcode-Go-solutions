package leetcode

import (
	"math"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return risValidBST(root, math.MinInt64, math.MaxInt64)
}

func risValidBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}
	if root.Val <= min || root.Val >= max {
		return false
	}
	return risValidBST(root.Left, min, root.Val) &&
		risValidBST(root.Right, root.Val, max)
}

func TestIsValidBST(t *testing.T) {
	output := isValidBST(&TreeNode{Val: 2,
		Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil}})

	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isValidBST(&TreeNode{Val: 5,
		Left: &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: &TreeNode{Val: 6, Left: nil, Right: nil}}})

	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
