package leetcode

import (
	"container/heap"
	"reflect"
	"testing"
)

type Word struct {
	word     string
	quantity int
}

type WordHeap []Word

func (w WordHeap) Len() int { return len(w) }

func (w WordHeap) Less(i, j int) bool {
	if w[i].quantity == w[j].quantity {
		return w[i].word < w[j].word
	}
	return w[i].quantity > w[j].quantity
}

func (w WordHeap) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func (w *WordHeap) Push(x interface{}) { *w = append(*w, x.(Word)) }

func (w *WordHeap) Pop() interface{} {
	x := (*w)[len(*w)-1]
	*w = (*w)[:len(*w)-1]
	return x
}

func topKFrequent(words []string, k int) []string {
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}
	w := WordHeap{}
	heap.Init(&w)
	for word, quantity := range counter {
		heap.Push(&w, Word{word: word, quantity: quantity})
	}
	result := make([]string, k)
	for i := 0; i < k; i++ {
		result[i] = heap.Pop(&w).(Word).word
	}
	return result
}

func TestTopKFrequent(t *testing.T) {
	output := topKFrequent([]string{"i", "love", "leetcode", "i", "love", "coding"}, 2)
	expected := []string{"i", "love"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
