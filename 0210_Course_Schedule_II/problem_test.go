package leetcode

import (
	"reflect"
	"testing"
)

type OrderGenerator struct {
	coursesToTake map[int][]int
	visited       []int
	result        []int
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	og := newOrderGenerator(numCourses, prerequisites)
	for i := 0; i < numCourses; i++ {
		if og.visited[i] == 0 && og.dfs(i) == false {
			return []int{}
		}
	}
	for i, j := 0, len(og.result)-1; i < j; i, j = i+1, j-1 {
		og.result[i], og.result[j] = og.result[j], og.result[i]
	}
	return og.result
}

func newOrderGenerator(numCourses int, prerequisites [][]int) *OrderGenerator {
	coursesToTake := map[int][]int{}
	for _, course := range prerequisites {
		coursesToTake[course[1]] = append(coursesToTake[course[1]], course[0])
	}
	return &OrderGenerator{
		coursesToTake: coursesToTake, visited: make([]int, numCourses), result: []int{}}
}

func (og *OrderGenerator) dfs(vertex int) bool {
	og.visited[vertex] = 1
	for _, course := range og.coursesToTake[vertex] {
		if og.visited[course] == 1 || (og.visited[course] == 0 && og.dfs(course) == false) {
			return false
		}
	}
	og.visited[vertex] = 2
	og.result = append(og.result, vertex)
	return true
}

func TestFindOrder(t *testing.T) {
	output := findOrder(2, [][]int{{1, 0}})
	expected := []int{0, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	expected = []int{0, 2, 1, 3}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findOrder(1, [][]int{})
	expected = []int{0}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
