package leetcode

import (
	"reflect"
	"strconv"
	"testing"
)

func compress(chars []byte) int {
	i, j := 0, 0
	for i < len(chars) {
		count := 1
		chars[j] = chars[i]
		i, j = i+1, j+1
		for ; i < len(chars) && chars[i-1] == chars[i]; i++ {
			count++
		}
		if count > 1 {
			str := strconv.Itoa(count)
			for _, c := range str {
				chars[j] = byte(c)
				j++
			}
		}
	}
	return j
}

func TestCompress(t *testing.T) {
	input := []byte("aabbccc")
	output := compress(input)
	expected := []byte("a2b2c3")
	if !reflect.DeepEqual(input[:output], expected) {
		t.Errorf("\noutput:   %c\nexpected: %c", input[:output], expected)
	}

	input = []byte("a")
	output = compress(input)
	expected = []byte("a")
	if !reflect.DeepEqual(input[:output], expected) {
		t.Errorf("\noutput:   %c\nexpected: %c", input[:output], expected)
	}

	input = []byte("abbbbbbbbbbbb")
	output = compress(input)
	expected = []byte("ab12")
	if !reflect.DeepEqual(input[:output], expected) {
		t.Errorf("\noutput:   %c\nexpected: %c", input[:output], expected)
	}

	input = []byte("aaabbaa")
	output = compress(input)
	expected = []byte("a3b2a2")
	if !reflect.DeepEqual(input[:output], expected) {
		t.Errorf("\noutput:   %c\nexpected: %c", input[:output], expected)
	}
}
