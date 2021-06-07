package leetcode

import "testing"

func reorganizeString(s string) string {
	letters := map[byte]int{}
	maxLenChar, maxLen := byte(0), 0
	for _, l := range s {
		letters[byte(l)]++
		if letters[byte(l)] > (len(s)+1)/2 {
			return ""
		}
		if letters[byte(l)] > maxLen {
			maxLenChar, maxLen = byte(l), letters[byte(l)]
		}
	}
	delete(letters, maxLenChar)
	var result = make([]byte, len(s))
	index := 0
	for ; maxLen > 0; maxLen-- {
		result[index] = maxLenChar
		index += 2
	}
	for char, value := range letters {
		for i := 0; i < value; i++ {
			if index >= len(s) {
				index = 1
			}
			result[index] = char
			index += 2
		}
	}
	return string(result)
}

func TestReorganizeString(t *testing.T) {
	output := reorganizeString("aab")
	expected := "aba"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = reorganizeString("aaab")
	expected = ""
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
