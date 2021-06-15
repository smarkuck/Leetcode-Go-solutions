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

func goodNodes(root *TreeNode) int {
	return travel(root, math.MinInt64)
}

func travel(node *TreeNode, prevMax int) int {
	if node == nil {
		return 0
	}
	result := 0
	if node.Val >= prevMax {
		prevMax = node.Val
		result++
	}
	return result + travel(node.Left, prevMax) + travel(node.Right, prevMax)
}

func TestGoodNodes(t *testing.T) {
	tree := &TreeNode{Val: 3,
		Left: &TreeNode{Val: 1,
			Left:  &TreeNode{Val: 3, Left: nil, Right: nil},
			Right: nil},
		Right: &TreeNode{Val: 4,
			Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil}}}

	output := goodNodes(tree)
	expected := 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	tree = &TreeNode{Val: 3,
		Left: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 4, Left: nil, Right: nil},
			Right: &TreeNode{Val: 2, Left: nil, Right: nil}},
		Right: nil}

	output = goodNodes(tree)
	expected = 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = goodNodes(&TreeNode{Val: 1, Left: nil, Right: nil})
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
