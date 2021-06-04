package leetcode

import (
	"reflect"
	"testing"
)

type Remover struct {
	input    string
	result   []byte
	validSet map[string]struct{}
}

func removeInvalidParentheses(s string) []string {
	return newRemover(s).remove()
}

func newRemover(input string) *Remover {
	return &Remover{input: input, result: []byte{}, validSet: make(map[string]struct{})}
}

func (r *Remover) remove() []string {
	redundantLeft, redundantRight := r.getRedundantParentheses()
	r.rremove(0, 0, 0, redundantLeft, redundantRight)

	result := []string{}
	for k := range r.validSet {
		result = append(result, k)
	}
	return result
}

func (r *Remover) getRedundantParentheses() (int, int) {
	remainingLeft, remainingRight := 0, 0
	for _, c := range r.input {
		if c == '(' {
			remainingLeft++
		} else if c == ')' {
			if remainingLeft == 0 {
				remainingRight++
			} else {
				remainingLeft--
			}
		}
	}
	return remainingLeft, remainingRight
}

func (r *Remover) rremove(i int, leftCount, rightCount int, remainingLeft, remainingRight int) {
	if i >= len(r.input) {
		if remainingLeft == 0 && remainingRight == 0 {
			r.validSet[string(r.result)] = struct{}{}
		}
		return
	}

	if r.input[i] == '(' && remainingLeft > 0 {
		r.rremove(i+1, leftCount, rightCount, remainingLeft-1, remainingRight)
	} else if r.input[i] == ')' && remainingRight > 0 {
		r.rremove(i+1, leftCount, rightCount, remainingLeft, remainingRight-1)
	}

	r.result = append(r.result, r.input[i])

	if r.input[i] == '(' {
		r.rremove(i+1, leftCount+1, rightCount, remainingLeft, remainingRight)
	} else if r.input[i] != ')' {
		r.rremove(i+1, leftCount, rightCount, remainingLeft, remainingRight)
	} else if rightCount < leftCount {
		r.rremove(i+1, leftCount, rightCount+1, remainingLeft, remainingRight)
	}

	r.result = r.result[:len(r.result)-1]
}

func TestRemoveInvalidParentheses(t *testing.T) {
	output := removeInvalidParentheses("()())()")
	expected := []string{"(())()", "()()()"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = removeInvalidParentheses("(a)())()")
	expected = []string{"(a())()", "(a)()()"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = removeInvalidParentheses(")(")
	expected = []string{""}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
