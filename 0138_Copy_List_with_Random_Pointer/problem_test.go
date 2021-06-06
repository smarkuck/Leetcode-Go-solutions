package leetcode

import (
	"reflect"
	"testing"
)

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	randomPointersMap := map[*Node]int{}
	resultPointers := []*Node{}
	result := &Node{Val: 0, Next: nil, Random: nil}
	for i, node, resultNode := 0, head, result; node != nil; i++ {
		resultNode.Next = &Node{Val: node.Val, Next: nil, Random: nil}
		randomPointersMap[node] = i
		resultPointers = append(resultPointers, resultNode.Next)
		node, resultNode = node.Next, resultNode.Next
	}
	for i, node, resultNode := 0, head, result.Next; node != nil; i++ {
		if i, ok := randomPointersMap[node.Random]; ok {
			resultNode.Random = resultPointers[i]
		}
		node, resultNode = node.Next, resultNode.Next
	}
	return result.Next
}

func TestCopyRandomList(t *testing.T) {
	list := &Node{Val: 7, Next: &Node{Val: 13, Next: &Node{Val: 11,
		Next: &Node{Val: 10, Next: &Node{Val: 1, Next: nil,
			Random: nil}, Random: nil}, Random: nil}, Random: nil}, Random: nil}
	list.Next.Random = list
	list.Next.Next.Random = list.Next.Next.Next.Next
	list.Next.Next.Next.Random = list.Next.Next
	list.Next.Next.Next.Next.Random = list

	output := copyRandomList(list)
	expected := list
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toArray(output), toArray(expected))
	}

	list = &Node{Val: 1, Next: &Node{Val: 2, Next: nil, Random: nil}, Random: nil}
	list.Random = list.Next
	list.Next.Random = list.Next

	output = copyRandomList(list)
	expected = list
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toArray(output), toArray(expected))
	}

	list = &Node{Val: 3, Next: &Node{Val: 3, Next: &Node{Val: 3, Next: nil,
		Random: nil}, Random: nil}, Random: nil}
	list.Next.Random = list

	output = copyRandomList(list)
	expected = list
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toArray(output), toArray(expected))
	}

	output = copyRandomList(nil)
	expected = nil
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toArray(output), toArray(expected))
	}
}

func toArray(n *Node) [][]int {
	a := [][]int{}
	for n != nil {
		r := []int{n.Val}
		if n.Random != nil {
			r = append(r, n.Random.Val)
		}
		a = append(a, r)
		n = n.Next
	}
	return a
}
