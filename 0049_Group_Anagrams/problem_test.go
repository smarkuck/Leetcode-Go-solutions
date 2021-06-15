package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	anagrams := map[[26]byte][]string{}
	for _, str := range strs {
		hash := [26]byte{}
		for _, c := range str {
			hash[c-'a']++
		}
		anagrams[hash] = append(anagrams[hash], str)
	}
	result := [][]string{}
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}

func TestGroupAnagrams(t *testing.T) {
	output := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	sort.Slice(output, func(i, j int) bool { return len(output[i]) < len(output[j]) })
	expected := [][]string{{"bat"}, {"tan", "nat"}, {"eat", "tea", "ate"}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = groupAnagrams([]string{""})
	expected = [][]string{{""}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
	output = groupAnagrams([]string{"a"})
	expected = [][]string{{"a"}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
