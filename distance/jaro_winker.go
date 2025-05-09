package distance

// Sources:
// https://github.com/jhvst/go-jaro-winkler-distance
// https://www.geeksforgeeks.org/jaro-and-jaro-winkler-similarity

// Bulk of the code is from this but with some optimizations and modern golang
// https://github.com/jhvst/go-jaro-winkler-distance

import (
	"math"
	"strings"
)

func JaroWinkerDistance(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}

	if len(s1) > len(s2) {
		temp := s1
		s1 = s2
		s2 = temp
	}

	runes1 := []byte(s1)
	runes2 := []byte(s2)
	l1 := len(runes1)
	l2 := len(runes2)

	window := int(math.Floor(float64(max(l1, l2))/2.0) - 1)

	m := 0
	// Since this is the max amount of matches
	//and we always allocate this amount of memory
	matches := make([]bool, l1)

	for i := range l1 {
		match := false
		if runes1[i] == runes2[i] {
			m++
			match = true
		}
		matches[i] = match
	}

	if m == 0 {
		return 0.0
	}

	t := 0.0
	slider := runes2[0:window]

	for i := range l1 {
		if matches[i] {
			continue
		}

		idx := strings.Index(string(slider), string(runes1[i]))
		if idx != -1 && !matches[idx] {
			t += 0.5
			matches[idx] = true
		}

		start := int(max(0, i-window))
		end := int(min(i+window, l1))

		slider_new := runes2[start:end]
		if len(slider_new) >= window {
			slider = slider_new
		}
	}

	term1 := float64(m) / float64(l1)
	term2 := float64(m) / float64(l2)
	term3 := (float64(m) - t) / float64(m)

	simj := (term1 + term2 + term3) / 3.0

	p := 0.1
	l := 0

	commonP := min(4, l1)

	s1p := s1[0:commonP]
	s2p := s2[0:commonP]
	for i := range s1p {
		if s1p[i] == s2p[i] {
			l++
		}
	}

	return simj + float64(l)*p*(1-simj)
}
