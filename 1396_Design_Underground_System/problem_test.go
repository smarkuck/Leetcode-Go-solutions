package leetcode

import "testing"

type Customer struct {
	startStation string
	startTime    int
}

type UndergroundSystem struct {
	customers    map[int]Customer
	averageTimes map[string]map[string][2]int
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{customers: map[int]Customer{}, averageTimes: map[string]map[string][2]int{}}
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.customers[id] = Customer{startStation: stationName, startTime: t}
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	customer := us.customers[id]
	delete(us.customers, id)
	averageTime := us.averageTimes[customer.startStation][stationName]
	averageTime[0] += t - customer.startTime
	averageTime[1]++
	if _, ok := us.averageTimes[customer.startStation]; !ok {
		us.averageTimes[customer.startStation] = map[string][2]int{stationName: averageTime}
	} else {
		us.averageTimes[customer.startStation][stationName] = averageTime
	}
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	averageTime := us.averageTimes[startStation][endStation]
	return float64(averageTime[0]) / float64(averageTime[1])
}

func TestUndergroundSystem(t *testing.T) {
	system := Constructor()
	system.CheckIn(45, "Leyton", 3)
	system.CheckIn(32, "Paradise", 8)
	system.CheckIn(27, "Leyton", 10)
	system.CheckOut(45, "Waterloo", 15)
	system.CheckOut(27, "Waterloo", 20)
	system.CheckOut(32, "Cambridge", 22)

	output := system.GetAverageTime("Paradise", "Cambridge")
	expected := 14.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = system.GetAverageTime("Leyton", "Waterloo")
	expected = 11.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	system.CheckIn(10, "Leyton", 24)
	output = system.GetAverageTime("Leyton", "Waterloo")
	expected = 11.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	system.CheckOut(10, "Waterloo", 38)
	output = system.GetAverageTime("Leyton", "Waterloo")
	expected = 12.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	system = Constructor()
	system.CheckIn(10, "Leyton", 3)
	system.CheckOut(10, "Paradise", 8)

	output = system.GetAverageTime("Leyton", "Paradise")
	expected = 5.0
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	system.CheckIn(5, "Leyton", 10)
	system.CheckOut(5, "Paradise", 16)
	output = system.GetAverageTime("Leyton", "Paradise")
	expected = 5.5
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	system.CheckIn(2, "Leyton", 21)
	system.CheckOut(2, "Paradise", 30)
	output = system.GetAverageTime("Leyton", "Paradise")
	expected = 6.666666666666667
	if output != expected {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
