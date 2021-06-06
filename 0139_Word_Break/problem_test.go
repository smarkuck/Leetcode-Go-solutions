package leetcode

import (
	"testing"
)

func wordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	for _, word := range wordDict {
		words[word] = true
	}
	dp := make(map[int]bool)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := range dp {
			if words[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func TestWordBreak(t *testing.T) {
	output := wordBreak("leetcode", []string{"leet", "code"})
	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = wordBreak("applepenapple", []string{"apple", "pen"})
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
