package leetcode

import (
	"reflect"
	"testing"
)

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	return clone(node, make(map[*Node]*Node))
}

func clone(node *Node, m map[*Node]*Node) *Node {
	if node == nil {
		return nil
	}
	if found, ok := m[node]; ok {
		return found
	}
	cloned := &Node{node.Val, nil}
	m[node] = cloned
	for _, neighbor := range node.Neighbors {
		cloned.Neighbors = append(cloned.Neighbors, clone(neighbor, m))
	}
	return cloned
}

func TestCloneGraph(t *testing.T) {
	node1 := &Node{Val: 1, Neighbors: nil}
	node2 := &Node{Val: 2, Neighbors: nil}
	node3 := &Node{Val: 3, Neighbors: nil}
	node4 := &Node{Val: 4, Neighbors: nil}

	node1.Neighbors = append(node1.Neighbors, node2, node4)
	node2.Neighbors = append(node2.Neighbors, node1, node3)
	node3.Neighbors = append(node3.Neighbors, node2, node4)
	node4.Neighbors = append(node4.Neighbors, node1, node3)

	output := cloneGraph(node1)
	expected := node1
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toMap(output), toMap(expected))
	}

	node1 = &Node{Val: 1, Neighbors: nil}
	output = cloneGraph(node1)
	expected = node1
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toMap(output), toMap(expected))
	}

	output = cloneGraph(nil)
	expected = nil
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toMap(output), toMap(expected))
	}

	node1 = &Node{Val: 1, Neighbors: nil}
	node2 = &Node{Val: 2, Neighbors: nil}

	node1.Neighbors = append(node1.Neighbors, node2)
	node2.Neighbors = append(node2.Neighbors, node1)

	output = cloneGraph(node1)
	expected = node1
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", toMap(output), toMap(expected))
	}
}

func toMap(node *Node) map[int][]int {
	result := map[int][]int{}
	visited := map[*Node]bool{}
	stack := []*Node{node}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if !visited[n] {
			visited[n] = true
			for _, neighbor := range n.Neighbors {
				result[n.Val] = append(result[n.Val], neighbor.Val)
				stack = append(stack, neighbor)
			}
		}
	}
	return result
}
