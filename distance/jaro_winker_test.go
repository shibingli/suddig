package distance_test

import (
	"testing"

	"github.com/VincentBrodin/suddig/distance"
)

// Taken from https://github.com/jhvst/go-jaro-winkler-distance/blob/master/algo_test.go
func TestJhvst(t *testing.T) {
	dist := distance.JaroWinkerDistance("accomodate", "accommodate")
	t.Logf("Expected 0.7509090909090909 got %f\n", dist)
	if dist != 0.7509090909090909 {
		t.FailNow()
	}

}

func TestJaroWinkerSameLength(t *testing.T) {
	dist := distance.JaroWinkerDistance("book", "back")
	t.Logf("Expected 0.667 got %f\n", dist)
	if dist != 0.667 {
		t.FailNow()
	}
}

func TestJaroWinkerTwoWords(t *testing.T) {
	dist := distance.JaroWinkerDistance("hello", "hello world")
	t.Logf("Expected 6 got %f\n", dist)
	if dist != 6 {
		t.FailNow()
	}
}

func TestJaroWinkerEmptyStrings(t *testing.T) {
	dist := distance.JaroWinkerDistance("", "")
	t.Logf("Expected 0 got %f\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestJaroWinkerBigStrings(t *testing.T) {
	a := string(make([]byte, 1000))
	b := string(make([]byte, 1000))
	dist := distance.JaroWinkerDistance(a, b)
	t.Logf("Expected 0 got %f\n", dist)
	if dist != 0 {
		t.FailNow()
	}
}

func TestJaroWinkerRepeating(t *testing.T) {
	dist := distance.JaroWinkerDistance("aabaa", "aaaa")
	t.Logf("Expected 1 got %f\n", dist)
	if dist != 1 {
		t.FailNow()
	}
}
