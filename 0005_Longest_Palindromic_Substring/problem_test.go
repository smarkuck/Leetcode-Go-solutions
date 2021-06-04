package leetcode

import (
	"strings"
	"testing"
)

func longestPalindrome(s string) string {
	s = strings.ReplaceAll(s, "", "|")
	palindromeRadii := make([]int, len(s))
	center, radius := 0, 0

	for center < len(s) {
		for center-radius-1 >= 0 && center+radius+1 < len(s) &&
			s[center-radius-1] == s[center+radius+1] {
			radius++
		}

		palindromeRadii[center] = radius
		oldCenter, oldRadius := center, radius
		center++
		radius = 0

		for center <= oldCenter+oldRadius {
			mirroredCenter := oldCenter - (center - oldCenter)
			maxMirroredRadius := oldCenter + oldRadius - center
			if palindromeRadii[mirroredCenter] < maxMirroredRadius {
				palindromeRadii[center] = palindromeRadii[mirroredCenter]
				center++
			} else if palindromeRadii[mirroredCenter] > maxMirroredRadius {
				palindromeRadii[center] = maxMirroredRadius
				center++
			} else {
				radius = maxMirroredRadius
				break
			}
		}
	}

	maxI, maxR := 0, 0
	for i, r := range palindromeRadii {
		if r > maxR {
			maxI, maxR = i, r
		}
	}

	return strings.ReplaceAll(s[maxI-maxR:maxI+maxR+1], "|", "")
}

func TestLongestPalindrome(t *testing.T) {
	output := longestPalindrome("babad")
	expected := "bab"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestPalindrome("cbbd")
	expected = "bb"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestPalindrome("a")
	expected = "a"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = longestPalindrome("ac")
	expected = "a"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
