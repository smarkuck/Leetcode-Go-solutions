package leetcode

import (
	"container/heap"
	"sort"
	"testing"
)

type Heap []int

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i] < h[j] }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	endTimes := &Heap{}
	for _, interval := range intervals {
		if endTimes.Len() > 0 && interval[0] >= (*endTimes)[0] {
			heap.Pop(endTimes)
		}
		heap.Push(endTimes, interval[1])
	}
	return endTimes.Len()
}

func TestMinMeetingRooms(t *testing.T) {
	output := minMeetingRooms([][]int{{0, 30}, {5, 10}, {15, 20}})
	expected := 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = minMeetingRooms([][]int{{7, 10}, {2, 4}})
	expected = 1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
