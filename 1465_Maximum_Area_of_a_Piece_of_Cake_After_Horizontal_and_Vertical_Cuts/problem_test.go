package leetcode

import (
	"math"
	"sort"
	"testing"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	return (findMaxGap(h, horizontalCuts) * findMaxGap(w, verticalCuts)) % (int(math.Pow10(9)) + 7)
}

func findMaxGap(size int, cuts []int) int {
	sort.Ints(cuts)
	maxGap := cuts[0]
	if size-cuts[len(cuts)-1] > maxGap {
		maxGap = size - cuts[len(cuts)-1]
	}
	for i := 1; i < len(cuts); i++ {
		if cuts[i]-cuts[i-1] > maxGap {
			maxGap = cuts[i] - cuts[i-1]
		}
	}
	return maxGap
}

func TestMaxArea(t *testing.T) {
	output := maxArea(5, 4, []int{1, 2, 4}, []int{1, 3})
	expected := 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = maxArea(5, 4, []int{3, 1}, []int{1})
	expected = 6
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = maxArea(5, 4, []int{3}, []int{3})
	expected = 9
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
