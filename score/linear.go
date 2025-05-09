package score

// Basic linear score function, that takes the distance divided by the max amount of changes needed.
// Great for Levenshtein Distance.
func Linear(query, target string, dist float64) float64 {
	maxDist := max(len(query), len(target))
	return 1.0 - dist/float64(maxDist)
}
