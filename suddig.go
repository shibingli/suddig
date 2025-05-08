package suddig

import "github.com/VincentBrodin/suddig/matcher"

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
