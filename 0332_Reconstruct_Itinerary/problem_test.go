package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

type PathFinder struct {
	tickets map[string][]string
	path    []string
	idx     int
}

func findItinerary(tickets [][]string) []string {
	return newPathFinder(tickets).findPath("JFK")
}

func newPathFinder(tickets [][]string) *PathFinder {
	pf := &PathFinder{tickets: map[string][]string{},
		path: make([]string, len(tickets)+1),
		idx:  len(tickets)}
	for _, ticket := range tickets {
		pf.tickets[ticket[0]] = append(pf.tickets[ticket[0]], ticket[1])
	}
	for _, v := range pf.tickets {
		sort.Strings(v)
	}
	return pf
}

func (pf *PathFinder) findPath(from string) []string {
	pf.rfindPath(from)
	return pf.path
}

func (pf *PathFinder) rfindPath(from string) {
	for len(pf.tickets[from]) > 0 {
		to := pf.tickets[from][0]
		pf.tickets[from] = pf.tickets[from][1:]
		pf.rfindPath(to)
	}
	pf.path[pf.idx] = from
	pf.idx--
}

func TestFindItinerary(t *testing.T) {
	output := findItinerary([][]string{
		{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}})
	expected := []string{"JFK", "MUC", "LHR", "SFO", "SJC"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = findItinerary([][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}})
	expected = []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
