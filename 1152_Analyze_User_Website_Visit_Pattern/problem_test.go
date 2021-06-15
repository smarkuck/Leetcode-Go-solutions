package leetcode

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

type SequenceData struct {
	visits   int
	lastUser string
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
	order := make([]int, len(username))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		i, j = order[i], order[j]
		if username[i] == username[j] {
			return timestamp[i] < timestamp[j]
		} else {
			return username[i] < username[j]
		}
	})
	threeSequences := make(map[string]SequenceData)
	maxThree, maxCount := "", 0
	for i := 0; i < len(username); i++ {
		currentUser := username[order[i]]
		for j := i + 1; j < len(username) && currentUser == username[order[j]]; j++ {
			for k := j + 1; k < len(username) && currentUser == username[order[k]]; k++ {
				seq := website[order[i]] + ":" + website[order[j]] + ":" + website[order[k]]
				if v, ok := threeSequences[seq]; !ok || v.lastUser != currentUser {
					v.lastUser = currentUser
					v.visits++
					threeSequences[seq] = v
					if v.visits > maxCount {
						maxThree, maxCount = seq, v.visits
					} else if v.visits == maxCount && seq < maxThree {
						maxThree = seq
					}
				}
			}
		}
	}
	return strings.Split(maxThree, ":")
}

func TestMostVisitedPattern(t *testing.T) {
	output := mostVisitedPattern(
		[]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"},
		[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		[]string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"})

	expected := []string{"home", "about", "career"}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
