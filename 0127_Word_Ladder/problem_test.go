package leetcode

import "testing"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]struct{}, len(wordList))
	for _, w := range wordList {
		wordSet[w] = struct{}{}
	}
	if _, ok := wordSet[endWord]; !ok {
		return 0
	}
	delete(wordSet, beginWord)
	delete(wordSet, endWord)

	beginSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	beginSet[beginWord], endSet[endWord] = struct{}{}, struct{}{}

	for length := 2; len(beginSet) > 0 && len(endSet) > 0; length++ {
		if len(beginSet) > len(endSet) {
			beginSet, endSet = endSet, beginSet
		}
		newBeginSet := make(map[string]struct{})
		for word := range beginSet {
			wordChars := []byte(word)
			for i, wordChar := range wordChars {
				for char := byte('a'); char <= 'z'; char++ {
					wordChars[i] = char
					newWord := string(wordChars)
					if _, ok := endSet[newWord]; ok {
						return length
					}
					if _, ok := wordSet[newWord]; ok {
						newBeginSet[newWord] = struct{}{}
						delete(wordSet, newWord)
					}
				}
				wordChars[i] = wordChar
			}
		}
		beginSet = newBeginSet
	}
	return 0
}

func TestLadderLength(t *testing.T) {
	output := ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"})
	expected := 5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log"})
	expected = 0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
