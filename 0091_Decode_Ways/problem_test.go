package leetcode

import "testing"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	t1, t2 := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i-1] == '0' || s[i-1] > '2' || (s[i-1] == '2' && s[i] > '6') {
			t1 = 0
		}
		if s[i] != '0' {
			t1 += t2
		}
		t1, t2 = t2, t1
	}
	return t2
}

func TestNumDecodings(t *testing.T) {
	output := numDecodings("12")
	expected := 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = numDecodings("226")
	expected = 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = numDecodings("0")
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = numDecodings("06")
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
