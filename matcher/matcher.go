package matcher

import "github.com/VincentBrodin/suddig/configs"

type Matcher struct {
	config configs.Config
}

func New(config configs.Config) *Matcher {
	return &Matcher{config: config}
}

func (m *Matcher) Match(query, target string) bool {
	a := m.config.StringFunc(query)
	b := m.config.StringFunc(target)
	dist := m.config.DistanceFunc(a, b)
	score := m.config.ScoreFunc(a, b, dist)
	return score >= m.config.MinScore
}

func (m *Matcher) Distance(query, target string) float64 {
	a := m.config.StringFunc(query)
	b := m.config.StringFunc(target)
	return m.config.DistanceFunc(a, b)
}

func (m *Matcher) Score(query, target string) float64 {
	a := m.config.StringFunc(query)
	b := m.config.StringFunc(target)
	dist := m.config.DistanceFunc(a, b)
	return m.config.ScoreFunc(a, b, dist)
}
