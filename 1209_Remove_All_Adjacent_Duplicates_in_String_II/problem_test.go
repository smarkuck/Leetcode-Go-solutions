package leetcode

import (
	"strings"
	"testing"
)

type Letter struct {
	letter rune
	count  int
}

func removeDuplicates(s string, k int) string {
	stack := []Letter{}
	for _, c := range s {
		if len(stack) == 0 || stack[len(stack)-1].letter != c {
			stack = append(stack, Letter{letter: c, count: 1})
		} else if stack[len(stack)-1].count == k-1 {
			stack = stack[:len(stack)-1]
		} else {
			stack[len(stack)-1].count++
		}
	}

	result := strings.Builder{}
	for _, l := range stack {
		result.WriteString(strings.Repeat(string(l.letter), l.count))
	}
	return result.String()
}

func TestRemoveDuplicates(t *testing.T) {
	output := removeDuplicates("abcd", 2)
	expected := "abcd"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = removeDuplicates("deeedbbcccbdaa", 3)
	expected = "aa"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = removeDuplicates("pbbcggttciiippooaais", 2)
	expected = "ps"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
