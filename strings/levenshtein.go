package strings

// Levenshtein calculates the Levenshtein distance between two input strings a
// and b and returns the distance as an integer.
//
// The Levenshtein distance algorithm measures the minimum number of
// single-character edits (insertions, deletions, or substitutions) required to
// transform one string into the other.
func Levenshtein(a, b string) int {
	la := len(a)
	lb := len(b)
	d := make([]int, la+1)

	// Initialize the first row of the matrix.
	for i := 1; i <= la; i++ {
		d[i] = i
	}

	var lastdiag, olddiag, temp int
	for i := 1; i <= lb; i++ {
		// Initialize the first element of the current row.
		d[0] = i
		lastdiag = i - 1
		for j := 1; j <= la; j++ {
			olddiag = d[j]

			// Calculate the minimum of three values:
			// - the value to the left plus 1
			// - the value above plus 1
			// - the diagonal value plus 0 or 1 (depending on whether the characters are equal)
			min := d[j] + 1
			if (d[j-1] + 1) < min {
				min = d[j-1] + 1
			}
			if (a)[j-1] == (b)[i-1] {
				temp = 0
			} else {
				temp = 1
			}
			if (lastdiag + temp) < min {
				min = lastdiag + temp
			}

			// Update the current element in the matrix.
			d[j] = min
			lastdiag = olddiag
			olddiag = min
		}
	}

	// The Levenshtein distance is the value in the last cell of the matrix.
	return d[la]
}
