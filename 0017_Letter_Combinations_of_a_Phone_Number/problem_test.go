package leetcode

import (
	"reflect"
	"testing"
)

var numberToLetters = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz"}

func letterCombinations(digits string) []string {
	result := []string{}
	if digits == "" {
		return result
	}
	generateStrings("", digits, &result)
	return result
}

func generateStrings(prefix string, pattern string, result *[]string) {
	if len(pattern) == 0 {
		*result = append(*result, prefix)
		return
	}
	for _, l := range numberToLetters[pattern[0]] {
		generateStrings(prefix+string(l), pattern[1:], result)
	}
}

func TestLetterCombinations(t *testing.T) {
	output := letterCombinations("23")
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = letterCombinations("")
	expected = []string{}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = letterCombinations("2")
	expected = []string{"a", "b", "c"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
