package leetcode

import (
	"strings"
	"testing"
)

func intToRoman(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var roman strings.Builder

	for i, value := range values {
		for num >= value {
			num -= value
			roman.WriteString(symbols[i])
		}
	}
	return roman.String()
}

func TestIntToRoman(t *testing.T) {
	output := intToRoman(3)
	expected := "III"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = intToRoman(4)
	expected = "IV"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = intToRoman(9)
	expected = "IX"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = intToRoman(58)
	expected = "LVIII"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = intToRoman(1994)
	expected = "MCMXCIV"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
