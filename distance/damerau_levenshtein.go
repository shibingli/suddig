package distance

// Sources:
// - https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance

// The Damerau–Levenshtein distance differs from the classical
// Levenshtein distance by including transpositions among its allowable operations
// in addition to the three classical single-character edit operations (insertions, deletions and substitutions).
// This is the Osa version of the damerau levenshtein distance
func DamerauLevenshteinDistance(s1, s2 string) float64 {
	// Early exit if the strings are the same
	if s1 == s2 {
		return 0
	}

	r1, r2 := []rune(s1), []rune(s2) // Use runes to solve problem with non ascii characters counting as 2 or more bytes

	// Create the matrix matrix
	matrix := make([][]int32, len(r1)+1)
	for i := range matrix {
		matrix[i] = make([]int32, len(r2)+1)
	}

	// Set prefixes AKA setting the edges
	for i := range max(len(r1), len(r2)) + 1 {
		if len(r1) >= i {
			matrix[i][0] = int32(i)
		}
		if len(r2) >= i {
			matrix[0][i] = int32(i)
		}
	}

	// Fill the matrix
	for j, jr := range r2 {
		for i, ir := range r1 {
			var subCost int32 = 1
			if jr == ir {
				subCost = 0
			}

			matrix[i+1][j+1] = min(
				matrix[i][j+1]+1,
				matrix[i+1][j]+1,
				matrix[i][j]+subCost)

			if i >= 1 && j >= 1 && r1[i] == r2[j-1] && r1[i-1] == r2[j] {
				matrix[i+1][j+1] = min(matrix[i+1][j+1], matrix[i-1][j-1]+1)
			}
		}
	}

	// Return the distance
	return float64(matrix[len(r1)][len(r2)])
}
