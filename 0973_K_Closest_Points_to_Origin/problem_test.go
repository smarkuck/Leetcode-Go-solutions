package leetcode

import (
	"reflect"
	"testing"
)

func kClosest(points [][]int, k int) [][]int {
	quickSort(points, 0, len(points)-1, k)
	result := [][]int{}
	for i := 0; i < k; i++ {
		result = append(result, points[i])
	}
	return result
}

func quickSort(a [][]int, l, r int, k int) {
	pivot := l
	for i := l; i < r; i++ {
		if a[i][0]*a[i][0]+a[i][1]*a[i][1] <= a[r][0]*a[r][0]+a[r][1]*a[r][1] {
			a[i], a[pivot] = a[pivot], a[i]
			pivot++
		}
	}
	a[pivot], a[r] = a[r], a[pivot]
	if pivot+1 == k {
		return
	} else if pivot+1 < k {
		quickSort(a, pivot+1, r, k)
	} else {
		quickSort(a, l, pivot-1, k)
	}
}

func TestKClosest(t *testing.T) {
	output := kClosest([][]int{{1, 3}, {-2, 2}}, 1)
	expected := [][]int{{-2, 2}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = kClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)
	expected = [][]int{{3, 3}, {-2, 4}}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
