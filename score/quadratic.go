package score

// Sharp drop near midpoint: scores around 0.5 shrink to 0.25 (0.5²),
// punishing mediocre matches, 
// while high-score plateau keeps ratios above ~0.9 near 0.81 (0.9²), 
// preserving top matches.
func Quadratic(query, target string, dist int32) float64 {
	maxDist := max(len(query), len(target))
	ratio := 1.0 - float64(dist)/float64(maxDist)
	return ratio * ratio
}
