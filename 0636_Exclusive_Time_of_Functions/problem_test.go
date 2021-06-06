package leetcode

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type TimePeriod struct {
	start    int
	duration int
}

func exclusiveTime(n int, logs []string) []int {
	times := make([]int, n)
	stack := []TimePeriod{}
	for _, log := range logs {
		logInfo := strings.Split(log, ":")
		if logInfo[1] == "start" {
			start, _ := strconv.Atoi(logInfo[2])
			stack = append(stack, TimePeriod{start: start, duration: 0})
		} else {
			timePeriod := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stop, _ := strconv.Atoi(logInfo[2])
			id, _ := strconv.Atoi(logInfo[0])
			time := stop - timePeriod.start + 1
			times[id] += time - timePeriod.duration
			if len(stack) > 0 {
				stack[len(stack)-1].duration += time
			}
		}
	}
	return times
}

func TestExclusiveTime(t *testing.T) {
	output := exclusiveTime(2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"})
	expected := []int{3, 4}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = exclusiveTime(1, []string{
		"0:start:0", "0:start:2", "0:end:5", "0:start:6", "0:end:6", "0:end:7"})
	expected = []int{8}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = exclusiveTime(2, []string{
		"0:start:0", "0:start:2", "0:end:5", "1:start:6", "1:end:6", "0:end:7"})
	expected = []int{7, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = exclusiveTime(2, []string{
		"0:start:0", "0:start:2", "0:end:5", "1:start:7", "1:end:7", "0:end:8"})
	expected = []int{8, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = exclusiveTime(1, []string{"0:start:0", "0:end:0"})
	expected = []int{1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
