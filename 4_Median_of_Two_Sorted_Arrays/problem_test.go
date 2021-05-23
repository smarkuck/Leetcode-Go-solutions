package leetcode

import (
	"math"
	"testing"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	allElements := len(nums1) + len(nums2)
	elementsLowerOrEqualMedian := (allElements + 1) / 2
	minNum1LeftPartSize, maxNum1LeftPartSize := 0, len(nums1)

	for minNum1LeftPartSize <= maxNum1LeftPartSize {
		num1LeftPartSize := (minNum1LeftPartSize + maxNum1LeftPartSize) / 2
		num2LeftPartSize := elementsLowerOrEqualMedian - num1LeftPartSize

		lastNum1LeftPartElement := getLastLeftPartElement(nums1, num1LeftPartSize)
		firstNum1RightPartElement := getFirstRightPartElement(nums1, num1LeftPartSize)
		lastNum2LeftPartElement := getLastLeftPartElement(nums2, num2LeftPartSize)
		firstNum2RightPartElement := getFirstRightPartElement(nums2, num2LeftPartSize)

		if lastNum1LeftPartElement <= firstNum2RightPartElement &&
			lastNum2LeftPartElement <= firstNum1RightPartElement {

			median := math.Max(
				float64(lastNum1LeftPartElement), float64(lastNum2LeftPartElement))

			if allElements%2 == 1 {
				return median
			} else {
				median2 := math.Min(
					float64(firstNum1RightPartElement), float64(firstNum2RightPartElement))
				return (median + median2) / 2
			}
		}

		if lastNum1LeftPartElement > firstNum2RightPartElement {
			maxNum1LeftPartSize = num1LeftPartSize - 1
		} else {
			minNum1LeftPartSize = num1LeftPartSize + 1
		}
	}
	return 0
}

func getLastLeftPartElement(nums []int, leftPartSize int) int {
	if leftPartSize-1 < 0 {
		return math.MinInt32
	}
	return nums[leftPartSize-1]
}

func getFirstRightPartElement(nums []int, leftPartSize int) int {
	if leftPartSize >= len(nums) {
		return math.MaxInt32
	}
	return nums[leftPartSize]
}

func TestFindMedianSortedArrays(t *testing.T) {
	output := findMedianSortedArrays([]int{1, 3}, []int{2})
	expected := 2.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	expected = 2.5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findMedianSortedArrays([]int{0, 0}, []int{0, 0})
	expected = 0.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findMedianSortedArrays([]int{}, []int{1})
	expected = 1.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findMedianSortedArrays([]int{2}, []int{})
	expected = 2.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
