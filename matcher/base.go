package matcher

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

func (m *Matcher) FindMatches(query string, targets []string) []string {
	arr := make([]string, len(targets))
	i := 0
	for _, s := range targets {
		if !m.Match(query, s) {
			continue
		}
		arr[i] = s
		i++
	}

	return arr[:i]
}

func (m *Matcher) RankMatches(query string, targets []string) []float64 {
	arr := make([]float64, len(targets))
	for i, s := range targets {
		arr[i] = m.Score(query, s)
	}

	return arr
}
