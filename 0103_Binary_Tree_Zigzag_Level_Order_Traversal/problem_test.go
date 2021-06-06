package leetcode

import (
	"reflect"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}
	prevQueue, queue := []*TreeNode{root}, []*TreeNode{}
	for swap := true; len(prevQueue) > 0; swap = !swap {
		level := []int{}
		for i := len(prevQueue) - 1; i >= 0; i-- {
			level = append(level, prevQueue[i].Val)
			if swap {
				if prevQueue[i].Left != nil {
					queue = append(queue, prevQueue[i].Left)
				}
				if prevQueue[i].Right != nil {
					queue = append(queue, prevQueue[i].Right)
				}
			} else {
				if prevQueue[i].Right != nil {
					queue = append(queue, prevQueue[i].Right)
				}
				if prevQueue[i].Left != nil {
					queue = append(queue, prevQueue[i].Left)
				}
			}
		}
		result = append(result, level)
		prevQueue = prevQueue[:0]
		prevQueue, queue = queue, prevQueue
	}
	return result
}

func TestZigzagLevelOrder(t *testing.T) {
	output := zigzagLevelOrder(&TreeNode{Val: 3,
		Left: &TreeNode{Val: 9, Left: nil, Right: nil},
		Right: &TreeNode{Val: 20,
			Left:  &TreeNode{Val: 15, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Left: nil, Right: nil}}})

	expected := [][]int{{3}, {20, 9}, {15, 7}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = zigzagLevelOrder(&TreeNode{Val: 1, Left: nil, Right: nil})
	expected = [][]int{{1}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = zigzagLevelOrder(nil)
	expected = [][]int{}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
