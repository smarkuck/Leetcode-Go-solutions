package leetcode

import (
	"math/rand"
	"testing"
)

type RandomizedSet struct {
	values  []int
	indexes map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{values: []int{}, indexes: map[int]int{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (r *RandomizedSet) Insert(val int) bool {
	if _, ok := r.indexes[val]; ok {
		return false
	}
	r.indexes[val] = len(r.values)
	r.values = append(r.values, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (r *RandomizedSet) Remove(val int) bool {
	if idx, ok := r.indexes[val]; !ok {
		return false
	} else {
		lastElement := r.values[len(r.values)-1]
		r.values[idx], r.indexes[lastElement] = lastElement, idx
		r.values = r.values[:len(r.values)-1]
		delete(r.indexes, val)
		return true
	}
}

/** Get a random element from the set. */
func (r *RandomizedSet) GetRandom() int {
	return r.values[rand.Intn(len(r.values))]
}

func TestRandomizedSet(t *testing.T) {
	set := Constructor()

	output := set.Insert(1)
	expected := true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = set.Remove(2)
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = set.Insert(2)
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output2 := set.GetRandom()
	expected2 := 2
	if output2 != expected2 {
		t.Errorf("\noutput:   %v\nexpected: %v", output2, expected2)
	}

	output = set.Remove(1)
	expected = true
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = set.Insert(2)
	expected = false
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output2 = set.GetRandom()
	expected2 = 2
	if output2 != expected2 {
		t.Errorf("\noutput:   %v\nexpected: %v", output2, expected2)
	}
}
