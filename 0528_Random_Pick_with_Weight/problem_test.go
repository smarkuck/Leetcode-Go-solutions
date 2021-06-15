package leetcode

import (
	"math/rand"
	"testing"
)

type Solution struct {
	sums []int
}

func Constructor(w []int) Solution {
	sums := []int{}
	prev := 0
	for _, value := range w {
		prev += value
		sums = append(sums, prev)
	}
	return Solution{sums: sums}
}

func (s *Solution) PickIndex() int {
	idx := rand.Intn(s.sums[len(s.sums)-1])
	l, r := 0, len(s.sums)-1
	for l <= r {
		m := (l + r) >> 1
		if idx < s.sums[m] {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func TestRandomPicker(t *testing.T) {
	picker := Constructor([]int{1})

	output := picker.PickIndex()
	expected := 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	picker = Constructor([]int{1, 3})

	output = picker.PickIndex()
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = picker.PickIndex()
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = picker.PickIndex()
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = picker.PickIndex()
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = picker.PickIndex()
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = picker.PickIndex()
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = picker.PickIndex()
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
