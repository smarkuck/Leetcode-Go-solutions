package leetcode

import "testing"

func lengthOfLongestSubstring(s string) int {
	maxLen, start := 0, 0
	table := [128]int{}
	for i := range table {
		table[i] = -1
	}
	for i, c := range s {
		if table[c] >= start {
			start = table[c] + 1
		}
		table[c] = i
		l := i - start + 1
		if l > maxLen {
			maxLen = l
		}
	}
	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	output := lengthOfLongestSubstring("abcabcbb")
	expected := 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = lengthOfLongestSubstring("bbbbb")
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = lengthOfLongestSubstring("pwwkew")
	expected = 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = lengthOfLongestSubstring("")
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
