package leetcode

import (
	"reflect"
	"testing"
)

func generateParenthesis(n int) []string {
	result := []string{}
	generator("", n, n, &result)
	return result
}

func generator(str string, l, r int, result *[]string) {
	if l == 0 && r == 0 {
		*result = append(*result, str)
	}
	if l > 0 {
		generator(str+"(", l-1, r, result)
	}
	if r > l {
		generator(str+")", l, r-1, result)
	}
}

func TestGenerateParenthesis(t *testing.T) {
	output := generateParenthesis(3)
	expected := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = generateParenthesis(1)
	expected = []string{"()"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
