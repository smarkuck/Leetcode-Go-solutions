package leetcode

import "testing"

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	buses := make(map[int]map[int]struct{})
	for i, route := range routes {
		buses[i] = make(map[int]struct{})
		for _, stop := range route {
			buses[i][stop] = struct{}{}
		}
	}

	stack := make([]map[int]struct{}, 0)
	for busId, route := range buses {
		if _, ok := route[source]; ok {
			if _, ok := route[target]; ok {
				return 1
			}
			stack = append(stack, route)
			delete(buses, busId)
		}
	}

	pathSize := 2
	for len(stack) > 0 {
		for busesLeft := len(stack); busesLeft > 0; busesLeft-- {
			currentBusRoute := stack[0]
			stack = stack[1:]
			for busId, busRoute := range buses {
				if hasConnection(currentBusRoute, busRoute) {
					if _, ok := busRoute[target]; ok {
						return pathSize
					}
					stack = append(stack, busRoute)
					delete(buses, busId)
				}
			}
		}
		pathSize++
	}

	return -1
}

func hasConnection(route1 map[int]struct{}, route2 map[int]struct{}) bool {
	if len(route1) > len(route2) {
		route1, route2 = route2, route1
	}

	for i := range route1 {
		if _, ok := route2[i]; ok {
			return true
		}
	}
	return false
}

func TestNumBusesToDestination(t *testing.T) {
	output := numBusesToDestination([][]int{{1, 2, 7}, {3, 6, 7}}, 1, 6)
	expected := 2
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = numBusesToDestination([][]int{{7, 12}, {4, 5, 15}, {6}, {15, 19}, {9, 12, 13}}, 15, 12)
	expected = -1
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
