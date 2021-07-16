package leetcode

import (
	"math"
	"testing"
)

type WordDistance struct {
	wordsMap map[string][]int
}

func Constructor(wordsDict []string) WordDistance {
	wd := WordDistance{wordsMap: map[string][]int{}}
	for i, word := range wordsDict {
		wd.wordsMap[word] = append(wd.wordsMap[word], i)
	}
	return wd
}

func (wd *WordDistance) Shortest(word1 string, word2 string) int {
	l1 := wd.wordsMap[word1]
	l2 := wd.wordsMap[word2]
	i1, i2, minGap := 0, 0, math.MaxInt64
	for i1 < len(l1) && i2 < len(l2) {
		gap := abs(l1[i1] - l2[i2])
		if gap < minGap {
			minGap = gap
		}
		if l1[i1] > l2[i2] {
			i2++
		} else {
			i1++
		}
	}
	return minGap
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func TestWordDistance(t *testing.T) {
	wd := Constructor([]string{"practice", "makes", "perfect", "coding", "makes"})

	output := wd.Shortest("coding", "practice")
	expected := 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = wd.Shortest("makes", "coding")
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
