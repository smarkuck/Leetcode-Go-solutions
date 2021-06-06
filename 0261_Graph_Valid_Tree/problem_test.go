package leetcode

import "testing"

func validTree(n int, edges [][]int) bool {
	if len(edges) != n-1 {
		return false
	}
	parents := make([]int, n)
	sizes := make([]int, n)
	for i := range parents {
		parents[i] = i
		sizes[i] = 1
	}
	find := func(n int) int {
		root := n
		for parents[root] != root {
			root = parents[root]
		}
		for n != root {
			parents[n], n = root, parents[n]
		}
		return root
	}
	union := func(n1, n2 int) bool {
		root1, root2 := find(n1), find(n2)
		if root1 == root2 {
			return false
		}
		if sizes[root1] < sizes[root2] {
			parents[root1] = root2
			sizes[root2] += sizes[root1]
		} else {
			parents[root2] = root1
			sizes[root1] += sizes[root2]
		}
		return true
	}
	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return false
		}
	}
	return true
}

func TestValidTree(t *testing.T) {
	output := validTree(5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}})
	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = validTree(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}})
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
