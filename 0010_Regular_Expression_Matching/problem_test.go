package leetcode

import "testing"

func isMatch(s string, p string) bool {
	n := len(s)
	m := len(p)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}

	dp[0][0] = true
	for column := 2; column <= m; column += 2 {
		if p[column-1] == '*' {
			dp[0][column] = dp[0][column-2]
		}
	}

	for row := 1; row <= n; row++ {
		for column := 1; column <= m; column++ {
			sc := s[row-1]
			pc := p[column-1]
			if sc == pc || pc == '.' {
				dp[row][column] = dp[row-1][column-1]
			} else if pc == '*' {
				if dp[row][column-2] {
					dp[row][column] = true
				} else if sc == p[column-2] || p[column-2] == '.' {
					dp[row][column] = dp[row-1][column]
				}
			}
		}
	}

	return dp[n][m]
}

func isMatch2(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	firstMatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
	if len(p) >= 2 && p[1] == '*' {
		return isMatch(s, p[2:]) || (firstMatch && isMatch(s[1:], p))
	}
	return firstMatch && isMatch(s[1:], p[1:])
}

func isMatch3(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if p[len(p)-1] == '*' {
		return isMatch(s, p[:len(p)-2]) ||
			(len(s) > 0) && (s[len(s)-1] == p[len(p)-2] || p[len(p)-2] == '.') &&
				isMatch(s[:len(s)-1], p)
	}
	return len(s) > 0 && (s[len(s)-1] == p[len(p)-1] || p[len(p)-1] == '.') &&
		isMatch(s[:len(s)-1], p[:len(p)-1])
}

func TestIsMatch(t *testing.T) {
	output := isMatch("aa", "a")
	expected := false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch("aa", "a*")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch("ab", ".*")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch("aab", "c*a*b")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch("mississippi", "mis*is*p*.")
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}

func TestIsMatch2(t *testing.T) {
	output := isMatch2("aa", "a")
	expected := false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch2("aa", "a*")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch2("ab", ".*")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch2("aab", "c*a*b")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch2("mississippi", "mis*is*p*.")
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}

func TestIsMatch3(t *testing.T) {
	output := isMatch3("aa", "a")
	expected := false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch3("aa", "a*")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch3("ab", ".*")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch3("aab", "c*a*b")
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = isMatch3("mississippi", "mis*is*p*.")
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
