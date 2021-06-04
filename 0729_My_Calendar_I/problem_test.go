package leetcode

import "testing"

type MyCalendar struct {
	root *Node
}

type Node struct {
	start, end  int
	left, right *Node
}

func Constructor() MyCalendar {
	return MyCalendar{root: nil}
}

func (mc *MyCalendar) Book(start int, end int) bool {
	if mc.root == nil {
		mc.root = &Node{start: start, end: end, left: nil, right: nil}
		return true
	}
	return mc.root.insert(&Node{start: start, end: end, left: nil, right: nil})
}

func (n *Node) insert(node *Node) bool {
	if node.start >= n.end {
		if n.right == nil {
			n.right = node
			return true
		}
		return n.right.insert(node)
	} else if node.end <= n.start {
		if n.left == nil {
			n.left = node
			return true
		}
		return n.left.insert(node)
	}
	return false
}

func TestCalendar(t *testing.T) {
	c := Constructor()
	output := c.Book(10, 20)
	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = c.Book(15, 25)
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = c.Book(20, 30)
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
