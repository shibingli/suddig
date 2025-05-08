package score

func Linear(query, target string, dist int32) float64 {
	maxDist := max(len(query), len(target))
	return 1.0 - float64(dist)/float64(maxDist)
}
