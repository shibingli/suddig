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
			buf := make([]float64, hi-lo)
			for i, s := range targets[lo:hi] {
				buf[i] = m.Score(query, s)
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
