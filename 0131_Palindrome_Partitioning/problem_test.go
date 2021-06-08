package leetcode

import (
	"reflect"
	"testing"
)

type PalindromeGenerator struct {
	s           string
	currentList []string
	dp          [][]bool
	result      [][]string
}

func partition(s string) [][]string {
	p := newPalindromeGenerator(s)
	p.dfs(0)
	return p.result
}

func newPalindromeGenerator(s string) *PalindromeGenerator {
	p := &PalindromeGenerator{
		s: s, currentList: []string{},
		dp: make([][]bool, len(s)), result: [][]string{}}
	for i := range p.dp {
		p.dp[i] = make([]bool, len(s))
	}
	return p
}

func (p *PalindromeGenerator) dfs(start int) {
	if start == len(p.s) {
		p.result = append(p.result, append([]string{}, p.currentList...))
	}
	for end := start; end < len(p.s); end++ {
		if p.s[start] == p.s[end] && (end-start <= 2 || p.dp[start+1][end-1]) {
			p.dp[start][end] = true
			p.currentList = append(p.currentList, p.s[start:end+1])
			p.dfs(end + 1)
			p.currentList = p.currentList[:len(p.currentList)-1]
		}
	}
}

func TestPartition(t *testing.T) {
	output := partition("aab")
	expected := [][]string{{"a", "a", "b"}, {"aa", "b"}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = partition("a")
	expected = [][]string{{"a"}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
