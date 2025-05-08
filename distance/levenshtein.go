package distance

// The Levenshtein algorithm returns a number based on how many edits is needed for the query to match the target.
// O(m*n) where m & n are the length of the input strigns
func LevenshteinDistance(query, target string) int32 {
	// Early exit if the strings are the same
	if query == target {
		return 0
	}

	// Create the matrix matrix
	matrix := make([][]int32, len(query)+1)
	for i := range matrix {
		matrix[i] = make([]int32, len(target)+1)
	}

	// Set prefixes AKA setting the edges
	for i := range max(len(query), len(target)) + 1 {
		if len(query) >= i {
			matrix[i][0] = int32(i)
		}
		if len(target) >= i {
			matrix[0][i] = int32(i)
		}
	}

	// Fill the matrix
	for j, jr := range target {
		for i, ir := range query {
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
	return matrix[len(query)][len(target)]
}
