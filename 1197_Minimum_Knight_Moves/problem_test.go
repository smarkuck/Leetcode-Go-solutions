package leetcode

import (
	"math"
	"testing"
)

func minKnightMoves(x int, y int) int {
	x, y = abs(x), abs(y)
	if x < y {
		x, y = y, x
	}
	if x == 1 && y == 0 {
		return 3
	}
	if x == 2 && y == 2 {
		return 4
	}
	delta := x - y
	if y > delta {
		return delta - 2*int(math.Floor(float64(delta-y)/3))
	} else {
		return delta - 2*((delta-y)/4)
	}
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func TestMinKnightMoves(t *testing.T) {
	output := minKnightMoves(2, 1)
	expected := 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = minKnightMoves(5, 5)
	expected = 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
