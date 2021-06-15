package leetcode

import "testing"

func findKthLargest(nums []int, k int) int {
	quickSelect(nums, 0, len(nums)-1, k-1)
	return nums[k-1]
}

func quickSelect(nums []int, l, r int, k int) {
	pivot, mid := l, (l+r)/2
	nums[r], nums[mid] = nums[mid], nums[r]
	for i := l; i < r; i++ {
		if nums[i] > nums[r] {
			nums[i], nums[pivot] = nums[pivot], nums[i]
			pivot++
		}
	}
	nums[r], nums[pivot] = nums[pivot], nums[r]
	if pivot == k {
		return
	} else if pivot > k {
		quickSelect(nums, l, pivot-1, k)
	} else {
		quickSelect(nums, pivot+1, r, k)
	}
}

func TestFindKthLargest(t *testing.T) {
	output := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	expected := 5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	expected = 4
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
