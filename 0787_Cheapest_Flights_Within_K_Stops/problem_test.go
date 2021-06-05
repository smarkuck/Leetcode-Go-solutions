package leetcode

import "testing"

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	inf := 1 << 31
	dPrev, d := make([]int, n), make([]int, n)
	for i := range d {
		dPrev[i], d[i] = inf, inf
	}
	dPrev[src], d[src] = 0, 0
	for i := 0; i < k+1; i++ {
		for _, flight := range flights {
			d[flight[1]] = min(d[flight[1]], dPrev[flight[0]]+flight[2])
		}
		dPrev, d = d, dPrev
	}
	if dPrev[dst] < inf {
		return dPrev[dst]
	}
	return -1
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestFindCheapestPrice(t *testing.T) {
	output := findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 1)
	expected := 200
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findCheapestPrice(3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}, 0, 2, 0)
	expected = 500
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
