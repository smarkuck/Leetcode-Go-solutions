package leetcode

import (
	"strconv"
	"testing"
)

func calculate(s string) int {
	cur, i := getNum(s, 0)
	last := 0
	for cur2, next := 0, 0; i < len(s); i++ {
		if s[i] == '+' {
			last += cur
			cur, next = getNum(s, i+1)
			i += next
		} else if s[i] == '-' {
			last += cur
			cur, next = getNum(s, i+1)
			cur *= -1
			i += next
		} else if s[i] == '*' {
			cur2, next = getNum(s, i+1)
			cur *= cur2
			i += next
		} else if s[i] == '/' {
			cur2, next = getNum(s, i+1)
			cur /= cur2
			i += next
		}
	}
	return last + cur
}

func getNum(s string, i int) (int, int) {
	start := i
	for ; s[i] < '0' || s[i] > '9'; i++ {
	}
	j := i
	for ; j < len(s) && s[j] >= '0' && s[j] <= '9'; j++ {
	}
	result, _ := strconv.Atoi(s[i:j])
	return result, j - start
}

func TestCalculate(t *testing.T) {
	output := calculate("3+2*2")
	expected := 7
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = calculate(" 3/2 ")
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = calculate(" 3+5 / 2 ")
	expected = 5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
