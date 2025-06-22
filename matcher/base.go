package matcher

func (m Matcher) Match(query, target string) bool {
	a := m.config.StringFunc(query)
	b := m.config.StringFunc(target)
	dist := m.config.DistanceFunc(a, b)
	score := m.config.ScoreFunc(a, b, dist)
	return score >= m.config.MinScore
}

func (m Matcher) Distance(query, target string) float64 {
	a := m.config.StringFunc(query)
	b := m.config.StringFunc(target)
	return m.config.DistanceFunc(a, b)
}

func (m Matcher) Score(query, target string) float64 {
	a := m.config.StringFunc(query)
	b := m.config.StringFunc(target)
	dist := m.config.DistanceFunc(a, b)
	return m.config.ScoreFunc(a, b, dist)
}

func (m Matcher) FindMatches(query string, targets []string) []string {
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

func (m Matcher) RankMatches(query string, targets []string) []float64 {
	arr := make([]float64, len(targets))
	for i, s := range targets {
		arr[i] = m.Score(query, s)
	}

	return arr
}

func (m Matcher) RankMatchesTokenized(query string, targets [][]string) []float64 {
	queryTokens := m.config.TokenFunc(query)
	size := float64(len(queryTokens))
	arr := make([]float64, len(targets))
	for i, tokens := range targets {
		avg := 0.0
		for _, q := range queryTokens {
			best := 0.0
			for _, token := range tokens {
				score := m.Score(q, token)
				if score > best {
					best = score
				}
			}
			avg += best
		}
		arr[i] = avg / size
	}
	return arr
}

func (m Matcher) TokenizeAll(targets []string) [][]string {
	arr := make([][]string, len(targets))
	for i, s := range targets {
		arr[i] = m.config.TokenFunc(s)
	}
	return arr
}
