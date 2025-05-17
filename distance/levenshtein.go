package distance

// Sources:
// - https://en.wikipedia.org/wiki/Levenshtein_distance

// The Levenshtein algorithm returns a number based on how many edits is needed for the query to match the target.
// O(m*n) where m & n are the length of the input strigns
// Levenshtein checks the three classical single-character edit operations (insertions, deletions and substitutions).
func LevenshteinDistance(s1, s2 string) float64 {
	// Early exit if the strings are the same
	if s1 == s2 {
		return 0
	}

	// Create the matrix matrix
	matrix := make([][]int32, len(s1)+1)
	for i := range matrix {
		matrix[i] = make([]int32, len(s2)+1)
	}

	// Set prefixes AKA setting the edges
	for i := range max(len(s1), len(s2)) + 1 {
		if len(s1) >= i {
			matrix[i][0] = int32(i)
		}
		if len(s2) >= i {
			matrix[0][i] = int32(i)
		}
	}

	// Fill the matrix
	for j, jr := range s2 {
		for i, ir := range s1 {
			var subCost int32 = 1
			if jr == ir {
				subCost = 0
			}

			matrix[i+1][j+1] = min(
				matrix[i][j+1]+1,
				matrix[i+1][j]+1,
				matrix[i][j]+subCost)
		}
	}

	// Return the distance
	return float64(matrix[len(s1)][len(s2)])
}
