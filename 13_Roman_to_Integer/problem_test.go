package leetcode

import "testing"

func romanToInt(s string) int {
	var value = map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	sum, last := 0, 0
	for _, c := range s {
		v := value[c]
		sum += v
		if v > last {
			sum -= 2 * last
		}
		last = v
	}
	return sum
}

func TestRomanToInt(t *testing.T) {
	output := romanToInt("III")
	expected := 3
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = romanToInt("IV")
	expected = 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = romanToInt("IX")
	expected = 9
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = romanToInt("LVIII")
	expected = 58
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = romanToInt("MCMXCIV")
	expected = 1994
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
