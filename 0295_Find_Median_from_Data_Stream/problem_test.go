package leetcode

import (
	"container/heap"
	"testing"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type MedianFinder struct {
	lo, hi *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	lo := &IntHeap{}
	hi := &IntHeap{}
	heap.Init(lo)
	heap.Init(hi)
	return MedianFinder{lo: lo, hi: hi}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.hi.Len() == 0 || num > (*mf.hi)[0] {
		heap.Push(mf.hi, num)
		if mf.hi.Len() > mf.lo.Len()+1 {
			heap.Push(mf.lo, -heap.Pop(mf.hi).(int))
		}
	} else {
		heap.Push(mf.lo, -num)
		if mf.lo.Len() > mf.hi.Len() {
			heap.Push(mf.hi, -heap.Pop(mf.lo).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	if mf.hi.Len() > mf.lo.Len() {
		return float64((*mf.hi)[0])
	}
	return float64((*mf.hi)[0]-(*mf.lo)[0]) / 2
}

func TestFindMedian(t *testing.T) {
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)

	output := mf.FindMedian()
	expected := 1.5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	mf.AddNum(3)
	output = mf.FindMedian()
	expected = 2.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
