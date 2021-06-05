package leetcode

import (
	"strings"
	"testing"
)

type WordBreaker struct {
	s        string
	wordDict []string
	itFailed map[int]bool
}

func newWordBreaker(s string, wordDict []string) *WordBreaker {
	return &WordBreaker{s: s, wordDict: wordDict, itFailed: map[int]bool{}}
}

func wordBreak(s string, wordDict []string) bool {
	return newWordBreaker(s, wordDict).resolve(0)
}

func (wb *WordBreaker) resolve(start int) bool {
	if start == len(wb.s) {
		return true
	}
	if wb.itFailed[start] {
		return false
	}
	for _, word := range wb.wordDict {
		if strings.HasPrefix(wb.s[start:], word) && wb.resolve(start+len(word)) {
			return true
		}
	}
	wb.itFailed[start] = true
	return false
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
