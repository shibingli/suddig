package suddig

import (
	"sort"

	"github.com/VincentBrodin/suddig/matcher"
)

// Match computes a normalized similarity score between two strings.
// It returns true if the similarity score is 80% or above.
// Match uses the base Distance function.
func Match(query, target string) bool {
	maxDist := max(len(query), len(target))
	percent := 1.0 - float64(Distance(query, target))/float64(maxDist)
	return percent >= 0.8
}

// Returns a normalized linear score based on distance.
// See Distance for the raw difference between the strings.
func Score(query, target string) float64 {
	m := matcher.New(matcher.DefualtConfig())
	return m.Score(query, target)
}

// Returns the distance based on the levenshtein distance.
// The Levenshtein algorithm returns a number based on how many edits is needed for the query to match the target.
// O(m*n) where m & n are the length of the input strigns.
// See Score for the percental difference between the strings.
func Distance(query, target string) int32 {
	m := matcher.New(matcher.DefualtConfig())
	return m.Distance(query, target)
}

// FindMatches returns all strings in targets whose similarity
// to query is ≥ 80%.
func FindMatches(query string, targets []string) []string {
	m := matcher.New(matcher.DefualtConfig())
	arr := make([]string, len(targets))
	i := 0
	for _, s := range targets {
		if m.Match(query, s) {
			arr[i] = s
			i++
		}
	}

	return arr[:i]
}

// RankMatches returns all strings in targets sorted in
// descending order of their similarity to query.
func RankMatches(query string, targets []string) []string {
	m := matcher.New(matcher.DefualtConfig())
	arr := make([]struct {
		s     string
		score float64
	}, len(targets))
	for i, s := range targets {
		arr[i] = struct {
			s     string
			score float64
		}{
			s:     s,
			score: m.Score(query, s),
		}

	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].score > arr[j].score
	})
	out := make([]string, len(targets))
	for i, v := range arr {
		out[i] = v.s
	}

	return out
}
