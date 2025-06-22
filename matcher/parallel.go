package matcher

import (
	"runtime"
	"sync"
)

// Identical to FindMatch execpt for the time complexity being O(n/wc)
// where n is the length of the targets and wc is the users computers cpu count
func (m *Matcher) ParallelMatch(query string, targets []string) []string {
	wc := runtime.NumCPU()
	n := len(targets)
	size := (n + wc - 1) / wc

	results := make([][]string, wc)

	var wg sync.WaitGroup
	wg.Add(wc)

	// Splits the work evenly between the workers
	for i := range wc {
		start := i * size
		end := min(start+size, n)

		go func(idx, lo, hi int) {
			defer wg.Done()
			if lo >= hi { // This solves the bug were the amount of items is less then the amount of workers making some workers empty
				results[idx] = []string{}
				return
			}
			buf := make([]string, 0, hi-lo)
			for _, s := range targets[lo:hi] {
				if m.Match(query, s) {
					buf = append(buf, s)
				}
			}
			results[idx] = buf
		}(i, start, end)
	}

	// wait for all workers
	wg.Wait()

	// flatten into single slice
	out := make([]string, 0, n)
	for _, buf := range results {
		out = append(out, buf...)
	}
	return out
}

// Identical to RankMatches execpt for the time complexity being O(n/wc)
// where n is the length of the targets and wc is the users computers cpu count
func (m *Matcher) ParallelRank(query string, targets []string) []float64 {
	wc := runtime.NumCPU()
	n := len(targets)
	size := (n + wc - 1) / wc

	results := make([][]float64, wc)

	var wg sync.WaitGroup
	wg.Add(wc)

	// Splits the work evenly between the workers
	for i := range wc {
		start := i * size
		end := min(start+size, n)

		go func(idx, lo, hi int) {
			defer wg.Done()
			if lo >= hi { // This solves the bug were the amount of items is less then the amount of workers making some workers empty
				results[idx] = []float64{}
				return
			}
			buf := make([]float64, hi-lo)
			for j, s := range targets[lo:hi] {
				buf[j] = m.Score(query, s)
			}
			results[idx] = buf
		}(i, start, end)
	}

	// wait for all workers
	wg.Wait()

	// flatten into single slice
	out := make([]float64, 0, n)
	for _, buf := range results {
		out = append(out, buf...)
	}
	return out
}

// Identical to RankMatchesTokenized execpt for the time complexity being O(n/wc)
// where n is the length of the targets and wc is the users computers cpu count
func (m *Matcher) ParallelRankTokenized(query string, targets [][]string) []float64 {
	wc := runtime.NumCPU()
	n := len(targets)
	size := (n + wc - 1) / wc

	results := make([][]float64, wc)

	var wg sync.WaitGroup
	wg.Add(wc)

	// Query tokens
	queryTokens := m.config.TokenFunc(query)
	count := float64(len(queryTokens))

	// Splits the work evenly between the workers
	for i := range wc {
		start := i * size
		end := min(start+size, n)

		go func(idx, lo, hi int) {
			defer wg.Done()
			if lo >= hi { // This solves the bug were the amount of items is less then the amount of workers making some workers empty
				results[idx] = []float64{}
				return
			}
			buf := make([]float64, hi-lo)
			for j, tokens := range targets[lo:hi] {
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
				buf[j] = avg / count
			}
			results[idx] = buf
		}(i, start, end)
	}

	// wait for all workers
	wg.Wait()

	// flatten into single slice
	out := make([]float64, 0, n)
	for _, buf := range results {
		out = append(out, buf...)
	}
	return out
}
