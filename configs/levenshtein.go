package configs

import (
	"strings"

	"github.com/VincentBrodin/suddig/distance"
	"github.com/VincentBrodin/suddig/matcher"
	"github.com/VincentBrodin/suddig/score"
)

func Levenshtein() matcher.Config {
	return matcher.Config{
		MinScore:     0.8,
		StringFunc:   strings.ToLower,
		DistanceFunc: distance.LevenshteinDistance,
		ScoreFunc:    score.Linear,
	}
}
