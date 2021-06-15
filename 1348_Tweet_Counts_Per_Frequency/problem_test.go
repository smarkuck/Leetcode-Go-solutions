package leetcode

import (
	"reflect"
	"testing"
)

var freqToSize = map[string]int{
	"minute": 60,
	"hour":   3600,
	"day":    86400,
}

type TweetCounts struct {
	tweets map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{tweets: map[string][]int{}}
}

func (tc *TweetCounts) RecordTweet(tweetName string, time int) {
	tc.tweets[tweetName] = append(tc.tweets[tweetName], time)
}

func (tc *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	chunkSize := freqToSize[freq]
	result := make([]int, (endTime-startTime)/chunkSize+1)
	for _, t := range tc.tweets[tweetName] {
		if t >= startTime && t <= endTime {
			result[(t-startTime)/chunkSize]++
		}
	}
	return result
}

func TestTweetCounter(t *testing.T) {
	counter := Constructor()
	counter.RecordTweet("tweet3", 0)
	counter.RecordTweet("tweet3", 60)
	counter.RecordTweet("tweet3", 10)

	output := counter.GetTweetCountsPerFrequency("minute", "tweet3", 0, 59)
	expected := []int{2}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	output = counter.GetTweetCountsPerFrequency("minute", "tweet3", 0, 60)
	expected = []int{2, 1}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}

	counter.RecordTweet("tweet3", 120)
	output = counter.GetTweetCountsPerFrequency("hour", "tweet3", 0, 210)
	expected = []int{4}
	if !reflect.DeepEqual(output, expected) {
		t.Errorf("\noutput:   %v\nexpected: %v", output, expected)
	}
}
