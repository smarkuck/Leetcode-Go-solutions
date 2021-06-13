package leetcode

import (
	"strings"
	"testing"
)

func minRemoveToMakeValid(s string) string {
	input := []byte(s)
	stack := []int{}
	for i, char := range input {
		if char == ')' {
			if len(stack) == 0 {
				input[i] = 0
			} else {
				stack = stack[:len(stack)-1]
			}
		} else if char == '(' {
			stack = append(stack, i)
		}
	}
	for _, i := range stack {
		input[i] = 0
	}
	return strings.ReplaceAll(string(input), "\000", "")
}

func TestMinRemoveToMakeValid(t *testing.T) {
	output := minRemoveToMakeValid("lee(t(c)o)de)")
	expected := "lee(t(c)o)de"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = minRemoveToMakeValid("a)b(c)d")
	expected = "ab(c)d"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = minRemoveToMakeValid("))((")
	expected = ""
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = minRemoveToMakeValid("(a(b(c)d)")
	expected = "a(b(c)d)"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
