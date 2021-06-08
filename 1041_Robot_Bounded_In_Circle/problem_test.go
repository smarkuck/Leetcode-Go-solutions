package leetcode

import "testing"

func isRobotBounded(instructions string) bool {
	x, y := 0, 0
	pos := 0
	directions := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for _, c := range instructions {
		switch c {
		case 'L':
			pos = (pos + 3) % 4
		case 'R':
			pos = (pos + 1) % 4
		case 'G':
			x += directions[pos][0]
			y += directions[pos][1]
		}
	}
	return pos != 0 || (x == 0 && y == 0)
}

func TestIsRobotBounded(t *testing.T) {
	output := isRobotBounded("GGLLGG")
	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isRobotBounded("GG")
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isRobotBounded("GL")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
