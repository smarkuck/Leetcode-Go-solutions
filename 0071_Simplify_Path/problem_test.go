package leetcode

import (
	"strings"
	"testing"
)

func simplifyPath(path string) string {
	words := strings.Split(path, "/")
	stack := []string{}
	for _, w := range words {
		if w == "" || w == "." {
			continue
		}
		if w == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}
		stack = append(stack, w)
	}
	return "/" + strings.Join(stack, "/")
}

func TestSimplifyPath(t *testing.T) {
	output := simplifyPath("/home/")
	expected := "/home"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = simplifyPath("/../")
	expected = "/"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = simplifyPath("/home//foo/")
	expected = "/home/foo"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = simplifyPath("/a/./b/../../c/")
	expected = "/c"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
