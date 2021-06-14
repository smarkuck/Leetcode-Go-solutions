package leetcode

import "testing"

type Value struct {
	value     string
	timestamp int
}

type TimeMap struct {
	values map[string][]Value
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{values: map[string][]Value{}}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.values[key] = append(t.values[key], Value{value: value, timestamp: timestamp})
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if v, ok := t.values[key]; !ok {
		return ""
	} else {
		i := findFirstLowerOrEqualIndex(v, timestamp)
		if i < 0 {
			return ""
		}
		return v[i].value
	}
}

func findFirstLowerOrEqualIndex(values []Value, timestamp int) int {
	l, r := 0, len(values)-1
	for l <= r {
		m := (l + r) >> 1
		if values[m].timestamp > timestamp {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return r
}

func TestTimeMap(t *testing.T) {
	timeMap := Constructor()
	timeMap.Set("foo", "bar", 1)

	output := timeMap.Get("foo", 1)
	expected := "bar"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = timeMap.Get("foo", 3)
	expected = "bar"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	timeMap.Set("foo", "bar2", 4)
	output = timeMap.Get("foo", 4)
	expected = "bar2"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = timeMap.Get("foo", 5)
	expected = "bar2"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	timeMap = Constructor()
	timeMap.Set("love", "high", 10)
	timeMap.Set("love", "low", 20)

	output = timeMap.Get("love", 5)
	expected = ""
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = timeMap.Get("love", 10)
	expected = "high"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = timeMap.Get("love", 15)
	expected = "high"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = timeMap.Get("love", 20)
	expected = "low"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = timeMap.Get("love", 25)
	expected = "low"
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
