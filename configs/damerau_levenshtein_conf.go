package configs

import (
	"strings"

	"github.com/VincentBrodin/suddig/distance"
	"github.com/VincentBrodin/suddig/score"
)

func DamerauLevenshtein() Config {
	return Config{
		MinScore:     0.8,
		StringFunc:   strings.ToLower,
		DistanceFunc: distance.DamerauLevenshteinDistance,
		ScoreFunc:    score.Linear,
	}
}
